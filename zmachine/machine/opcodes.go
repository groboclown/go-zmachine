// V1 compatible op-codes
package machine

import "fmt"

type extOpCode struct {
	OpCode
	form OpCodeForm
	op   uint8
	vers []int
}

var allOpCodes = []extOpCode{

	// ==========================================================================
	// Long Form OpCodes

	{
		form: LongOpCodeForm,
		op:   0x01,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "je",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x02,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "jl",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x03,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "jg",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x04,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "dec_chk",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x05,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "inc_chk",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x06,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "jin",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x07,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "test",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x08,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "or",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x09,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "and",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0a,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "test_attr",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0b,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "set_attr",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0c,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "clear_attr",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0d,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "store",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0e,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "insert_obj",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x0f,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "loadw",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x10,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "loadb",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x11,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_prop",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x12,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_prop_addr",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x13,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_next_prop",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x14,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "add",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x15,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "sub",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x16,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "mul",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x17,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "div",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x18,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "mod",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x19,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_2s",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x1a,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_2n",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x1b,
		vers: []int{5},
		OpCode: OpCode{
			Name:     "set_colour",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x1b,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "set_colour",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: LongOpCodeForm,
		op:   0x1c,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "throw",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	// ==========================================================================
	// Short Form OpCodes

	{
		form: ShortOpCodeForm,
		op:   0x00,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "jz",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x01,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_sibling",
			Stores:   true,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x02,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_child",
			Stores:   true,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x03,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_parent",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x04,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_prop_len",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x05,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "inc",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x06,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "dec",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x07,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_addr",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x08,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_1s",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x09,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "remove_obj",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0a,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_obj",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0b,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "ret",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0c,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "jump",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0d,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_paddr",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0e,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "load",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0f,
		vers: []int{1, 2, 3, 4},
		OpCode: OpCode{
			Name:     "not",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x0f,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_1n",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x30,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "rtrue",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x31,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "rfalse",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x32,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print",
			Stores:   false,
			Branches: false,
			Text:     true,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x33,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_ret",
			Stores:   false,
			Branches: false,
			Text:     true,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x34,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "nop",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x35,
		vers: []int{1, 2, 3},
		OpCode: OpCode{
			Name:     "v1_save",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x35,
		vers: []int{4},
		OpCode: OpCode{
			Name:     "v4_save",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	// save for v5+ is illegal.

	{
		form: ShortOpCodeForm,
		op:   0x36,
		vers: []int{1, 2, 3},
		OpCode: OpCode{
			Name:     "v1_restore",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x36,
		vers: []int{4},
		OpCode: OpCode{
			Name:     "v4_restore",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	// restore v5+ is illegal

	{
		form: ShortOpCodeForm,
		op:   0x37,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "restart",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x38,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "ret_popped",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x39,
		vers: []int{1, 2, 3, 4},
		OpCode: OpCode{
			Name:     "pop",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x39,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "catch",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x3a,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "quit",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x3b,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "new_line",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x3c,
		vers: []int{3}, // only valid for version 3.
		OpCode: OpCode{
			Name:     "show_status",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ShortOpCodeForm,
		op:   0x3d,
		vers: []int{3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "verify",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	// 0x3e is the first byte of the extended opcode.

	{
		form: ShortOpCodeForm,
		op:   0x3f,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "piracy",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	// ==========================================================================
	// Variable Form OpCodes

	// V1-3 form of the subroutine call.
	{
		form: VariableOpCodeForm,
		op:   0x00,
		vers: []int{1, 2, 3},
		OpCode: OpCode{
			Name:     "call",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	// V4+ form of the subroutine call.
	{
		form: VariableOpCodeForm,
		op:   0x00,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_vs",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x01,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "storew",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x02,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "storeb",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x03,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "put_prop",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x04,
		vers: []int{1, 2, 3},
		OpCode: OpCode{
			Name:     "v1_sread",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x04,
		vers: []int{4},
		OpCode: OpCode{
			Name:     "v4_sread",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x04,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "aread",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x05,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_char",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x06,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_num",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x07,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "random",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x08,
		vers: []int{1, 2, 3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "push",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x09,
		vers: []int{1, 2, 3, 4, 5},
		OpCode: OpCode{
			Name:     "v1_pull",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x09,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "v6_pull",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0a,
		vers: []int{3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "split_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0b,
		vers: []int{3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "set_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	// Secretly a double-variable op-code
	{
		form: VariableOpCodeForm,
		op:   0x0c,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_vs2",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0d,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "erase_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0e,
		vers: []int{4, 5},
		OpCode: OpCode{
			Name:     "v4_erase_line",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0e,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "v6_erase_line",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0f,
		vers: []int{4, 5},
		OpCode: OpCode{
			Name:     "v4_set_cursor",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x0f,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "v6_set_cursor",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x10,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "get_cursor",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x11,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "set_text_style",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x12,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "buffer_mode",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x13,
		vers: []int{3, 4},
		OpCode: OpCode{
			Name:     "v3_output_stream",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x13,
		vers: []int{5},
		OpCode: OpCode{
			Name:     "v5_output_stream",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x13,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "v6_output_stream",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x14,
		vers: []int{3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "input_stream",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x15,
		vers: []int{3, 4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "sound_effect",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x16,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "read_char",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x17,
		vers: []int{4, 5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "scan_table",
			Stores:   true,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x18,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "not",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x19,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_vn",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	// Secretly a double-variable op-code.
	{
		form: VariableOpCodeForm,
		op:   0x1a,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "call_vn2",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x1b,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "tokenize",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x1c,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "encode_text",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x1d,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "copy_table",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x1e,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_table",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: VariableOpCodeForm,
		op:   0x1f,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "check_arg_count",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	// ==========================================================================
	// Extended Form OpCodes

	{
		form: ExtendedOpCodeForm,
		op:   0x00,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "v5_save",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x01,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "v5_restore",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x02,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "log_shift",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x03,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "art_shift",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x04,
		vers: []int{5},
		OpCode: OpCode{
			Name:     "v5_set_font",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x04,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "v6_set_font",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x05,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "draw_picture",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x06,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "picture_data",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x07,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "erase_picture",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x08,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "set_margins",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x09,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "save_undo",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x0a,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "restore_undo",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x0b,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "print_unicode",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x0c,
		vers: []int{5, 6, 7, 8},
		OpCode: OpCode{
			Name:     "check_unicode",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x0d,
		vers: []int{5},
		OpCode: OpCode{
			Name:     "set_true_colour_v5",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x0d,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "set_true_colour_v6",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x10,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "move_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x11,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "window_size",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x12,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "window_style",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x13,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "get_wind_prop",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x14,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "scroll_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x15,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "pop_stack",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x16,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "read_mouse",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x17,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "mouse_window",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x18,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "push_stack",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x1a,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "print_form",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x1b,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "make_menu",
			Stores:   false,
			Branches: true,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x1c,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "picture_table",
			Stores:   false,
			Branches: false,
			Text:     false,
		},
	},

	{
		form: ExtendedOpCodeForm,
		op:   0x1d,
		vers: []int{6, 7, 8},
		OpCode: OpCode{
			Name:     "buffer_screen",
			Stores:   true,
			Branches: false,
			Text:     false,
		},
	},
}

// ==========================================================================

type splitOpCodes struct {
	variableOpcodes  map[uint8]*OpCode
	shortOpcodes     map[uint8]*OpCode
	longOpCodes      map[uint8]*OpCode
	doubleVarOpCodes map[uint8]*OpCode
	extendedOpCodes  map[uint8]*OpCode
}

type OpCodeForm int

const (
	VariableOpCodeForm OpCodeForm = iota
	ShortOpCodeForm
	LongOpCodeForm
	ExtendedOpCodeForm
)

// assembleOpCodes puts the op codes into appropriate maps.
func assembleOpCodes(version int) splitOpCodes {
	ret := splitOpCodes{
		variableOpcodes:  make(map[uint8]*OpCode),
		shortOpcodes:     make(map[uint8]*OpCode),
		longOpCodes:      make(map[uint8]*OpCode),
		doubleVarOpCodes: make(map[uint8]*OpCode),
		extendedOpCodes:  make(map[uint8]*OpCode),
	}

	for _, o := range allOpCodes {
		// Is this an opcode for the version?
		ok := false
		for _, v := range o.vers {
			if v == version {
				ok = true
				break
			}
		}
		if !ok {
			continue
		}

		// Put the opcode in the right field.
		switch o.form {
		case VariableOpCodeForm:
			// Hard-coded double var opcodes.
			//   Note that "26" is 0x1a, and "12" is 0x0c, and those identically matches the opcode id.
			//   These also have a variable type opcode.
			if o.op == 12 || o.op == 26 {
				setOpCode(ret.doubleVarOpCodes, o.op, &o.OpCode)
			} else {
				setOpCode(ret.variableOpcodes, o.op, &o.OpCode)
			}
		case ShortOpCodeForm:
			setOpCode(ret.shortOpcodes, o.op, &o.OpCode)
		case LongOpCodeForm:
			setOpCode(ret.longOpCodes, o.op, &o.OpCode)
		case ExtendedOpCodeForm:
			setOpCode(ret.extendedOpCodes, o.op, &o.OpCode)
		default:
			// Programmer error.
			panic("bad opcode form")
		}
	}

	return ret
}

func setOpCode(m map[uint8]*OpCode, op uint8, code *OpCode) {
	if v, ok := m[op]; ok {
		panic(fmt.Sprintf("already set op code %x: %s vs %s", op, v.Name, code.Name))
	}
	m[op] = code
}
