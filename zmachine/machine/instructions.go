// Instruction decoding.
package machine

import "fmt"

type OperandType int

const (
	LargeOperand    OperandType = 0b00
	SmallOperand    OperandType = 0b01
	VariableOperand OperandType = 0b10
	OmittedOperand  OperandType = 0b11
)

type OpCode struct {
	Name     string
	Stores   bool
	Branches bool
	Text     bool // 'true' only for 'print' and 'print_ret' instructions.
}

type OperandValueType int

const (
	ConstantByteValue OperandValueType = iota
	ConstantWordValue
	TopStackValue
	LocalVariableValue
	GlobalVariableValue
)

type Operand struct {
	Type  OperandValueType
	Value uint16
}

type Branch struct {
	BranchCondition bool // True means branch if true, False means branch if false
	Address         AbsAddr
	ReturnTrue      bool // Offset 0 means return false
	ReturnFalse     bool // Offset 1 means return true
}

type Instruction struct {
	Instruction *OpCode
	EndPos      AbsAddr // The last decoded byte address; the next instruction or piece of information to decode is at this + 1.
	Operands    []Operand
	Stores      *Operand // Variable to store
	Branch      *Branch
	Text        []rune
}

type OpDecode interface {
	DecodeAt(mem MemoryData, pos AbsAddr) (*Instruction, error)
}

type OpDecodeV1_4 struct {
	variableOpcodes  map[uint8]*OpCode
	shortOpcodes     map[uint8]*OpCode
	longOpCodes      map[uint8]*OpCode
	doubleVarOpCodes map[uint8]*OpCode
	zscii            Zscii
}

func (o OpDecodeV1_4) DecodeAt(mem MemoryData, pos AbsAddr) (*Instruction, error) {
	opcode := mem.ByteAt(pos)
	// Basic hard-coded opcode forms
	switch opcode {
	case 12:
		// call_vs2
		// Note that this is v4+ only, and call_vn2 (26) is 5+ only.
		// 2nd byte of operand types is given (at most 8 operands)
		return DecodeVariableDoubleForm(o.doubleVarOpCodes, opcode, mem, pos+1, o.zscii)
	default:
		// Decode based on the form
		rawForm := (opcode >> 6) & 0x3
		switch rawForm {
		case 0b11:
			return DecodeVariableForm(o.variableOpcodes, opcode, mem, pos+1, o.zscii)
		case 0b10:
			return DecodeShortForm(o.shortOpcodes, opcode, mem, pos+1, o.zscii)
		default:
			return DecodeLongForm(o.longOpCodes, opcode, mem, pos+1, o.zscii)
		}
	}
}

type OpDecodeV5_plus struct {
	OpDecodeV1_4
	extendedOpcodes map[uint8]*OpCode
}

func (o OpDecodeV5_plus) DecodeAt(mem MemoryData, pos AbsAddr) (*Instruction, error) {
	opcode := mem.ByteAt(pos)
	// Basic hard-coded opcode forms
	switch opcode {
	case 12:
		// call_vs2
		fallthrough
	case 26:
		// call_vn2
		// 2nd byte of operand types is given.
		return DecodeVariableDoubleForm(o.doubleVarOpCodes, opcode, mem, pos+1, o.zscii)
	case 0xbe:
		return DecodeExtendedForm(o.extendedOpcodes, mem, pos+1, o.zscii)
	default:
		// Decode based on the form
		rawForm := (opcode >> 6) & 0x3
		switch rawForm {
		case 0b11:
			return DecodeVariableForm(o.variableOpcodes, opcode, mem, pos+1, o.zscii)
		case 0b10:
			return DecodeShortForm(o.shortOpcodes, opcode, mem, pos+1, o.zscii)
		default:
			return DecodeLongForm(o.longOpCodes, opcode, mem, pos+1, o.zscii)
		}
	}
}

func DecodeShortForm(lookup map[uint8]*OpCode, opcode uint8, mem MemoryData, pos AbsAddr, zscii Zscii) (*Instruction, error) {
	// Bits 4 & 5 give the operand type.
	// Remember bit numbers: 76543210
	opType := OperandType((opcode >> 4) & 0b11)

	// OpCode Instruction itself is bottom 4 bits
	opCodeId := opcode & 0b1111
	opInstruction, err := LookupOpCode(lookup, opCodeId, "short")
	if err != nil {
		return nil, err
	}

	operands := make([]Operand, 0)
	if opType != OmittedOperand {
		// 1 operand
		o, p, err := DecodeOperand(opType, mem, pos+1)
		if err != nil {
			return nil, err
		}
		operands = append(operands, *o)
		pos = p
	}

	return DecodeStoreBranchText(
		opInstruction,
		operands,
		mem,
		pos,
		zscii,
	)
}

func DecodeLongForm(lookup map[uint8]*OpCode, opcode uint8, mem MemoryData, pos AbsAddr, zscii Zscii) (*Instruction, error) {
	// always 2 operands

	// bit 6 == first operand type
	op0Type := VariableOperand
	if opcode&0b1000000 == 0 {
		op0Type = SmallOperand
	}

	// bit 5 == second operand type
	op1Type := VariableOperand
	if opcode&0b100000 == 0 {
		op1Type = SmallOperand
	}

	// OpCode Instruction itself is bottom 5 bits
	opCodeId := opcode & 0b11111
	opInstruction, err := LookupOpCode(lookup, opCodeId, "long")
	if err != nil {
		return nil, err
	}

	op0, pos, err := DecodeOperand(op0Type, mem, pos)
	if err != nil {
		return nil, err
	}
	op1, pos, err := DecodeOperand(op1Type, mem, pos)
	if err != nil {
		return nil, err
	}

	return DecodeStoreBranchText(
		opInstruction,
		[]Operand{*op0, *op1},
		mem,
		pos,
		zscii,
	)
}

func DecodeVariableForm(lookup map[uint8]*OpCode, opcode uint8, mem MemoryData, pos AbsAddr, zscii Zscii) (*Instruction, error) {
	// OpCode Instruction is bottom 5 bits
	opInst, err := LookupOpCode(lookup, opcode&0b11111, "variable")
	if err != nil {
		return nil, err
	}

	pos++
	operandTypes := DecodeVarType(mem.ByteAt(pos))

	// If bit 5 == 0, opcode count is 2, if 1 then VAR.
	if opcode&0b100000 == 0 {
		// opcode count is 2
		operandTypes[2] = OmittedOperand
		operandTypes[3] = OmittedOperand
	}

	operands := make([]Operand, 0, 4)
	for _, ot := range operandTypes {
		if ot == OmittedOperand {
			break
		}
		o, p, err := DecodeOperand(ot, mem, pos)
		if err != nil {
			return nil, err
		}
		operands = append(operands, *o)
		pos = p
	}

	return DecodeStoreBranchText(
		opInst,
		operands,
		mem,
		pos,
		zscii,
	)
}

func DecodeExtendedForm(lookup map[uint8]*OpCode, mem MemoryData, pos AbsAddr, zscii Zscii) (*Instruction, error) {
	pos++
	opInstruction, err := LookupOpCode(lookup, mem.ByteAt(pos), "extended")
	if err != nil {
		return nil, err
	}

	// VAR opcode count; 0 - 4
	pos++
	operandTypes := DecodeVarType(mem.ByteAt(pos))

	operands := make([]Operand, 0, 4)
	for _, ot := range operandTypes {
		if ot == OmittedOperand {
			break
		}
		o, p, err := DecodeOperand(ot, mem, pos)
		if err != nil {
			return nil, err
		}
		operands = append(operands, *o)
		pos = p
	}

	return DecodeStoreBranchText(
		opInstruction,
		operands,
		mem,
		pos,
		zscii,
	)
}

func DecodeVariableDoubleForm(lookup map[uint8]*OpCode, opcode uint8, mem MemoryData, pos AbsAddr, zscii Zscii) (*Instruction, error) {
	opInst, err := LookupOpCode(lookup, opcode, "double-variable")
	if err != nil {
		return nil, err
	}
	pos++
	ops0 := DecodeVarType(mem.ByteAt(pos))
	pos++
	ops1 := DecodeVarType(mem.ByteAt(pos))
	operands := make([]Operand, 0)
	for _, ot := range []OperandType{ops0[0], ops0[1], ops0[2], ops0[3], ops1[0], ops1[1], ops1[2], ops1[3]} {
		if ot == OmittedOperand {
			// Always means the last operand, even if other items are not this.
			break
		}
		o, p, err := DecodeOperand(ot, mem, pos)
		if err != nil {
			return nil, err
		}
		operands = append(operands, *o)
		pos = p
	}

	return DecodeStoreBranchText(
		opInst,
		operands,
		mem,
		pos,
		zscii,
	)
}

func DecodeVarType(val uint8) [4]OperandType {
	ret := [4]OperandType{OmittedOperand, OmittedOperand, OmittedOperand, OmittedOperand}
	for i := 0; i < 4; i++ {
		t := OperandType((val & 0b11000000) >> 6)
		if t == OmittedOperand {
			return ret
		}
		ret[i] = t
		val = val << 2
	}
	return ret
}

// DecodeOperand decodes the operand type, whose value is at the given position.
//
// Returns the decoded operand, and the address of the last position read.
func DecodeOperand(opType OperandType, mem MemoryData, pos AbsAddr) (*Operand, AbsAddr, error) {
	switch opType {
	case LargeOperand:
		return &Operand{
			Type:  ConstantWordValue,
			Value: asWord(mem.ByteAt(pos), mem.ByteAt(pos+1)),
		}, pos + 1, nil
	case SmallOperand:
		return &Operand{
			Type:  ConstantByteValue,
			Value: uint16(mem.ByteAt(pos)),
		}, pos, nil
	case VariableOperand:
		val := mem.ByteAt(pos)
		if val == 0 {
			return &Operand{
				Type: TopStackValue,
			}, pos, nil
		} else if val <= 0xf {
			return &Operand{
				Type:  LocalVariableValue,
				Value: uint16(val - 1),
			}, pos, nil
		}
		return &Operand{
			Type:  GlobalVariableValue,
			Value: uint16(val - 0x10),
		}, pos, nil
	default:
		// Should have already ignored the Omitted type.
		return nil, 0, fmt.Errorf("invalid operand type %x", opType)
	}
}

func LookupOpCode(lookup map[uint8]*OpCode, opCodeId uint8, form string) (*OpCode, error) {
	opInstruction, ok := lookup[opCodeId]
	if !ok {
		return nil, fmt.Errorf("unknown %s-form opcode %x", form, opCodeId)
	}
	return opInstruction, nil
}

func DecodeStoreBranchText(
	instruction *OpCode,
	operands []Operand,
	mem MemoryData,
	pos AbsAddr,
	zscii Zscii,
) (*Instruction, error) {
	var stores *Operand = nil
	var branch *Branch = nil
	var text []rune = nil

	if instruction.Stores {
		s, p, err := DecodeOperand(VariableOperand, mem, pos+1)
		if err != nil {
			return nil, err
		}
		stores = s
		pos = p
	}
	if instruction.Branches {
		pos++
		var offset int16 = 0
		bOp := mem.ByteAt(pos)
		// Bit 7 == 1 then branch on true, else branch on false.
		branchOn := (bOp & 0b10000000) != 0
		if bOp&0b1000000 == 0 {
			// bit 6 == 0 then remainder + next byte is a 14-bit signed integer.
			pos++
			offset = asSigned14Bit(bOp, mem.ByteAt(pos))
		} else {
			offset = int16(bOp) & 0b00111111
		}
		// Branch-to address is computed as:
		//   Address after branch data + Offset - 2
		branch = &Branch{
			BranchCondition: branchOn,
			Address:         AbsAddr(pos + 1 + AbsAddr(offset) - 2),
			ReturnTrue:      offset == 0,
			ReturnFalse:     offset == 1,
		}
	}
	if instruction.Text {
		t, p, err := zscii.DecodeString(mem, pos, int(mem.Size()))
		if err != nil {
			return nil, err
		}
		text = t
		pos = p - 1
	}

	return &Instruction{
		Instruction: instruction,
		Operands:    operands,
		EndPos:      pos,
		Stores:      stores,
		Branch:      branch,
		Text:        text,
	}, nil
}
