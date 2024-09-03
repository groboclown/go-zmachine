// V2 story file specific information.
package machine

func Version2(mem VolatileMemoryData) (Version, error) {
	// Find the 32 abbreviations
	abbrevZ, err := NewZsciiV2(make([][]ZsciiChar, 0))
	if err != nil {
		return nil, err
	}
	abbrev, err := DecodeAbbreviationsTable(
		mem,
		asJoinedByteAddress(mem.ByteAt(0x18), mem.ByteAt(0x19)),
		32,
		abbrevZ,
	)
	if err != nil {
		return nil, err
	}

	z, err := NewZsciiV2(abbrev)
	if err != nil {
		return nil, err
	}
	ops := assembleOpCodes(2)
	return &v2Version{
		mem:    mem,
		header: &v2Header{v1Header: v1Header{mem: mem, marked: mem}},
		ops: &OpDecodeV1_4{
			variableOpcodes:  ops.variableOpcodes,
			shortOpcodes:     ops.shortOpcodes,
			longOpCodes:      ops.longOpCodes,
			doubleVarOpCodes: ops.doubleVarOpCodes,
			zscii:            z,
		},
	}, nil
}

type v2Version struct {
	mem    VolatileMemoryData
	header *v2Header
	ops    OpDecode
}

func (v *v2Version) Header() Header {
	return v.header
}

func (v *v2Version) Opcodes() OpDecode {
	return v.ops
}

func (v *v2Version) InitialRoutineState() *RoutineCallState {
	return v1InitialRoutineState(v.mem)
}

type v2Header struct {
	v1Header // V2 is nearly identical in structure to v1.
}

func (v *v2Header) AbbreviationsTableAddress() AbsAddr {
	return asByteAddress(asWord(v.v1Header.mem.ByteAt(0x18), v.v1Header.mem.ByteAt(0x19)))
}
