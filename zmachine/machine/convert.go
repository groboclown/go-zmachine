// Standard data conversion helpers.
package machine

// asWord converts 2 bytes into an unsigned 16 bit word.
//
// All integers are stored as big-endian.
func asWord(byte0, byte1 uint8) uint16 {
	w0 := uint16(byte0) & 0xff
	w1 := uint16(byte1) & 0xff
	return (w0 << 8) | w1
}

// fromWord turns an unsigned 16 bit word into its component 2 bytes.
//
// The returned bytes are the order they should be stored into memory.
func fromWord(word uint16) (uint8, uint8) {
	b1 := word & 0xff
	b0 := (word >> 8) & 0xff
	return uint8(b0), uint8(b1)
}

// asLong turns 4 bytes, in the order they occur in memory, into a 32-bit unsigned long.
func asLong(byte0, byte1, byte2, byte3 uint8) uint32 {
	l0 := uint32(byte0) & 0xff
	l1 := uint32(byte1) & 0xff
	l2 := uint32(byte2) & 0xff
	l3 := uint32(byte3) & 0xff
	return (l0 << 24) | (l1 << 16) | (l2 << 8) | l3
}

func fromLong(long uint32) (uint8, uint8, uint8, uint8) {
	b3 := long & 0xff
	b2 := (long >> 8) & 0xff
	b1 := (long >> 16) & 0xff
	b0 := (long >> 24) & 0xff
	return uint8(b0), uint8(b1), uint8(b2), uint8(b3)
}

// asByteAddress returns the address for the given byte pointer.
//
// "A byte address specifies a byte in memory in the range 0 up to the last byte of static memory."
func asByteAddress(pos uint16) AbsAddr {
	return AbsAddr(pos)
}

// asByteAddress returns the address for the given byte pointer.
//
// "A byte address specifies a byte in memory in the range 0 up to the last byte of static memory."
func asJoinedByteAddress(p0 uint8, p1 uint8) AbsAddr {
	return asByteAddress(asWord(p0, p1))
}

// asWordAddress returns the address for the given word pointer.
//
// "A word address specifies an even address in the bottom 128K of memory (by giving the address
// divided by 2). (Word addresses are used only in the abbreviations table.)"
func asWordAddress(pos uint16) AbsAddr {
	return AbsAddr(uint32(pos) * 2)
}

// asPackedRoutineAddress returns the address for the given packed routine pointer.
// Offset must be either the routine offset (header $28) or string offset (header $2a)
//
// "A packed address specifies where a routine or string begins in high memory."
func asPackedAddress(pos uint32, offset uint32, packedPtrMult uint16, packedOffsetMult uint16) AbsAddr {
	return AbsAddr((pos*uint32(packedPtrMult) + (offset * uint32(packedOffsetMult))))
}

// asSignedWord converts the unsigned word to a signed word.
func asSignedWord(word uint16) int16 {
	if word > 32767 {
		return int16(int32(word) - 65536)
	}
	return int16(int32(word) & 0x7fff)
}

// normalizeSignedWord returns a signed word into a normalized 16 bit word, suitable for memory storage.
func normalizeSignedWord(signed int16) uint16 {
	if signed < 0 {
		return uint16(65536+int32(signed)) & 0xffff
	}
	return uint16(int32(signed)) & 0x7fff
}

// asSigned14Bit converts the 2 bytes into a 14-bit signed number.
//
// This is used by the branch offset calculation.  The top 2 bits of the top byte are ignored.
func asSigned14Bit(top, bottom uint8) int16 {
	val := int16(top&0b00111111) | int16(bottom)
	if val > maxSigned14Bit {
		val = maxSigned14Bit - val
	}
	return val
}

const maxSigned14Bit = 0b0010000000000000
