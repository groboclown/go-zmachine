// The Z-Machine version information.
//
// The version heavily impacts the operation of the machine itself.
package machine

import "fmt"

type Version interface {
	Header() Header
	Opcodes() OpDecode
	InitialRoutineState() *RoutineCallState
}

func NewVersion(memory VolatileMemoryData) (Version, error) {
	// Version number is always the byte at offset 0.
	number := memory.ByteAt(0x00)

	switch number {
	case 1:
		return Version1(memory), nil
	case 2:
		return Version2(memory)
	case 3:
		return Version3(memory)
	case 4:
		return Version4(memory)
	case 5:
		return Version5(memory)
	case 6:
		return Version6(memory)
	case 7:
		return Version7(memory)
	case 8:
		return Version8(memory)
	default:
		return nil, fmt.Errorf("unsupported version number %d", number)
	}
}

/*
func Version1() *Version {
	return &Version{
		number:           1,
		maxStoryLength:   128,
		packedPtrMult:    2,
		packedOffsetMult: 0,
	}
}

func Version2() *Version {
	return &Version{
		number:           2,
		maxStoryLength:   128,
		packedPtrMult:    2,
		packedOffsetMult: 0,
	}
}

func Version3() *Version {
	return &Version{
		number:           3,
		maxStoryLength:   128,
		packedPtrMult:    2,
		packedOffsetMult: 0,
	}
}

func Version4() *Version {
	return &Version{
		number:           4,
		maxStoryLength:   256,
		packedPtrMult:    4,
		packedOffsetMult: 0,
	}
}

func Version5() *Version {
	return &Version{
		number:           5,
		maxStoryLength:   256,
		packedPtrMult:    4,
		packedOffsetMult: 0,
	}
}

func Version6() *Version {
	return &Version{
		number:           6,
		maxStoryLength:   512,
		packedPtrMult:    4,
		packedOffsetMult: 8,
	}
}

func Version7() *Version {
	return &Version{
		number:           7,
		maxStoryLength:   320,
		packedPtrMult:    4,
		packedOffsetMult: 8,
	}
}

func Version8() *Version {
	return &Version{
		number:           8,
		maxStoryLength:   512,
		packedPtrMult:    8,
		packedOffsetMult: 0,
	}
}

func (v *Version) Number() int {
	return v.number
}
*/
