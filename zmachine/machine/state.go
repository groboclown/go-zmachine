// Game state for iterating through the program.
package machine

// VMState contains the state of the
type VMState struct {
	Memory       VolatileMemoryData
	RoutineStack []*RoutineCallState
}

// Clone creates a copy of the VMState instance, for use in debugging or UNDO support.
func (v *VMState) Clone() *VMState {
	stack := make([]*RoutineCallState, len(v.RoutineStack))
	for i, r := range v.RoutineStack {
		stack[i] = r.Clone()
	}
	return &VMState{
		Memory:       v.Memory.Clone(),
		RoutineStack: stack,
	}
}

// RoutineCallState keeps track of a single routine's local variables and stack and position in execution.
type RoutineCallState struct {
	ProgramCounter AbsAddr
	Locals         []uint16
	Stack          []uint16
}

const MaxStackSize = 65535

// NewRoutineCallState creates a new routine call state for a routine with the initialized local variables.
//
// The programCounter must point to the first opcode in the routine.
func NewRoutineCallState(
	programCounter AbsAddr,
	locals []uint16,
) *RoutineCallState {
	vlocal := make([]uint16, len(locals))
	copy(vlocal, locals)
	return &RoutineCallState{
		ProgramCounter: programCounter,
		Locals:         vlocal,
		Stack:          make([]uint16, 0, MaxStackSize),
	}
}

func (r *RoutineCallState) Clone() *RoutineCallState {
	locals := make([]uint16, len(r.Locals))
	copy(locals, r.Locals)
	stack := make([]uint16, len(r.Stack))
	copy(stack, r.Stack)
	return &RoutineCallState{
		ProgramCounter: r.ProgramCounter,
		Locals:         locals,
		Stack:          stack,
	}
}
