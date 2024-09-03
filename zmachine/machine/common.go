// Common data types, shared between the different systems.
package machine

// AbsAddr points to an absolute position in the machine's memory.
type AbsAddr uint32

// MemoryRange marks a range of memory addresses.
type MemoryRange struct {
	Start AbsAddr // Memory starting at this address.
	End   AbsAddr // Final address position, inclusive.
}
