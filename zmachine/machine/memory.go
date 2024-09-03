// The memory state.
//
// The memory works by having a base, read-only version representing the
// raw game file, then a series of overlays for the VM state and the
// game modified memory.
//
// Changing some memory values act as an I/O to the machine state itself.
package machine

import "fmt"

type Memory struct {
	storyData          []uint8 // The base, read-only memory store.
	maxMemorySize      uint32
	dynamicMemoryBound AbsAddr
	wrData             map[AbsAddr]uint8 // Overwritten data
}

func NewMemory(story []uint8) *Memory {
	return &Memory{
		storyData:          story,
		maxMemorySize:      uint32(len(story)),
		dynamicMemoryBound: AbsAddr(len(story)), // temporary setting.
		wrData:             make(map[AbsAddr]uint8),
	}
}

func (m *Memory) Reset() {
	m.wrData = make(map[AbsAddr]uint8)
}

// Version discovers the story file version.
func (m *Memory) VersionNumber() int {
	// Always byte 0.
	return int(m.storyData[0])
}

// StoryData returns the original game's story data.
func (m *Memory) StoryData() MemoryData {
	return &rawMemory{m.storyData}
}

// Size returns the maximum size allowed for the memory.
func (m *Memory) Size() uint32 {
	return m.maxMemorySize
}

// SetDynamicMemoryBoundary sets the position of the dynamic memory location.
//
// The game may write only to dynamic memory and select areas of the header.
func (m *Memory) SetDynamicMemoryBoundary(pos AbsAddr) error {
	if pos >= AbsAddr(m.maxMemorySize) {
		return fmt.Errorf("dynamic memory bound (%d) outside maximum memory size (%d)", pos, m.maxMemorySize)
	}
	m.dynamicMemoryBound = pos
	return nil
}

func (m *Memory) ByteAt(pos AbsAddr) uint8 {
	v, ok := m.wrData[pos]
	if ok {
		return v
	}
	p := int(pos)
	if p < len(m.storyData) {
		return m.storyData[p]
	}
	return 0
}

func (m *Memory) WriteByteAt(value uint8, pos AbsAddr) error {
	if pos >= m.dynamicMemoryBound {
		return fmt.Errorf("illegal memory write position: %x", pos)
	}
	m.wrData[pos] = value
	return nil
}

func (m *Memory) Clone() VolatileMemoryData {
	writableData := make(map[AbsAddr]uint8)
	for k, v := range m.wrData {
		writableData[k] = v
	}
	return &Memory{
		storyData:          m.storyData, // read-only, so it's fine.
		maxMemorySize:      m.maxMemorySize,
		dynamicMemoryBound: m.dynamicMemoryBound,
		wrData:             writableData,
	}
}

type rawMemory struct {
	data []uint8
}

func (r *rawMemory) Size() uint32 {
	return uint32(len(r.data))
}

func (r *rawMemory) ByteAt(pos AbsAddr) uint8 {
	p := int(pos)
	if p < len(r.data) {
		return r.data[p]
	}
	return 0
}

// MemoryData provides an abstract way to get data from the machine memory.
type MemoryData interface {
	ByteAt(pos AbsAddr) uint8
	Size() uint32
}

// VolatileMemoryData allows the game to attempt to write to a memory position.
type VolatileMemoryData interface {
	MemoryData
	WriteByteAt(value uint8, pos AbsAddr) error

	// Clone allows for copying the data to aid in UNDO support.
	Clone() VolatileMemoryData
}

func IsMemoryBitSet(mem MemoryData, addr AbsAddr, bit int) bool {
	mask := uint8(1 << bit)
	return mem.ByteAt(addr)&mask != 0
}
