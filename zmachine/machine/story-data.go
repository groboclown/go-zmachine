package machine

import (
	"fmt"
)

// StoryData represents the raw story data.
//
// It also allows for parsing the header data, critical for setting up the initial machine state.
type StoryData struct {
	data []uint8

	// Parsed header data
	dynamicMemoryEnd AbsAddr
	staticMemoryEnd  AbsAddr
	highMemoryStart  AbsAddr
	stringOffset     uint16
	routineOffset    uint16
}

func NewStoryData(data []uint8) (*StoryData, error) {
	if len(data) < 16 {
		return nil, fmt.Errorf("story data must have at least 16 bytes")
	}
	dynamicMemoryEnd := asWord(data[headerDynamicMemoryEndAddr], data[headerDynamicMemoryEndAddr+1]) - 1
	staticMemoryEnd := len(data)
	if staticMemoryEnd > 0xffff {
		staticMemoryEnd = 0xffff
	}
	highMemoryStart := asWord(data[headerHighMemoryStartAddr], data[headerHighMemoryStartAddr+1])
	routineOffset := asWord(data[headerRoutineOffsetAddr], data[headerRoutineOffsetAddr+1])
	stringOffset := asWord(data[headerStringOffsetAddr], data[headerStringOffsetAddr+1])
	return &StoryData{
		data:             data,
		dynamicMemoryEnd: AbsAddr(dynamicMemoryEnd),
		staticMemoryEnd:  AbsAddr(staticMemoryEnd),
		highMemoryStart:  AbsAddr(highMemoryStart),
		routineOffset:    routineOffset,
		stringOffset:     stringOffset,
	}, nil
}

func (s *StoryData) DynamicMemoryRange() MemoryRange {
	return MemoryRange{
		Start: 0,
		End:   s.dynamicMemoryEnd,
	}
}

func (s *StoryData) StaticMemoryRange() MemoryRange {
	return MemoryRange{
		Start: s.dynamicMemoryEnd + 1,
		End:   s.staticMemoryEnd,
	}
}

func (s *StoryData) HighMemoryRange() MemoryRange {
	return MemoryRange{
		Start: s.highMemoryStart,
		End:   AbsAddr(len(s.data)),
	}
}

const (
	headerEndAddr              = 64
	headerHighMemoryStartAddr  = 0x04
	headerDynamicMemoryEndAddr = 0x0e
	headerRoutineOffsetAddr    = 0x28
	headerStringOffsetAddr     = 0x2a
)

// GetByte fetches an 8-bit unsigned integer from the given absolute address in the story data.
func (s *StoryData) GetByte(pos AbsAddr) (uint8, error) {
	if int(pos) >= len(s.data) {
		return 0, fmt.Errorf("story address out of range: %x", pos)
	}
	return s.data[pos], nil
}

// GetWord fetches a 16-bit unsigned integer from the given absolute address in the story data.
func (s *StoryData) GetWord(pos AbsAddr) (uint16, error) {
	if int(pos)+1 >= len(s.data) {
		return 0, fmt.Errorf("story address out of range: %x", pos)
	}
	return asWord(s.data[pos], s.data[pos+1]), nil
}

// GetWord fetches a 32-bit unsigned integer from the given absolute address in the story data.
func (s *StoryData) GetLong(pos AbsAddr) (uint32, error) {
	if int(pos)+3 >= len(s.data) {
		return 0, fmt.Errorf("story address out of range: %x", pos)
	}
	return asLong(s.data[pos], s.data[pos+1], s.data[pos+2], s.data[pos+3]), nil
}
