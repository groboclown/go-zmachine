// Arithmetic operations.
package machine

import (
	"errors"
)

// ArithmeticCompare performs the comparison between the two values, returning -1 if word0 is less, 0 if they are equal, 1 if word0 is greater.
//
// This implies a signed comparison.
func ArithmeticCompare(word0 uint16, word1 uint16) int {
	w0 := asSignedWord(word0)
	w1 := asSignedWord(word1)
	if w0 < w1 {
		return -1
	}
	if w0 == w1 {
		return 0
	}
	return 1
}

// ArithmeticOverflow handles converting a signed operation back into a memory word value.
func ArithmeticOverflow(word int32) uint16 {
	return normalizeSignedWord(int16(word & 0xffff))
}

// ArithmeticAdd adds the two words together, returning a signed 32-bit value.
//
// This implies a signed operation.  Overflow must be handled by the caller.
func ArithmeticAdd(word0 uint16, word1 uint16) int32 {
	return int32(asSignedWord(word0)) + int32(asSignedWord(word1))
}

// ArithmeticSubtract subtracts the two words together, returning a signed 32-bit value.
//
// This implies a signed operation.  Overflow must be handled by the caller.
func ArithmeticSubtract(word0 uint16, word1 uint16) int32 {
	return int32(asSignedWord(word0)) - int32(asSignedWord(word1))
}

// ArithmeticMultiply multiplies the two words together, returning a signed 32-bit value.
//
// This implies a signed operation.  Overflow must be handled by the caller.
func ArithmeticMultiply(word0 uint16, word1 uint16) int32 {
	return int32(asSignedWord(word0)) * int32(asSignedWord(word1))
}

// ArithmeticDivide divides the two words, returning a signed 32-bit value.
//
// This implies a signed operation.  Overflow must be handled by the caller.
func ArithmeticDivide(word0 uint16, word1 uint16) (int32, error) {
	div := int32(asSignedWord(word1))
	if div == 0 {
		return 0, errors.New("divide by zero")
	}
	return int32(asSignedWord(word0)) / div, nil
}

// ArithmeticRemainder returns the remainder-after-divide for the two words, returning a signed 32-bit value.
//
// This implies a signed operation.  Overflow must be handled by the caller.
func ArithmeticRemainder(word0 uint16, word1 uint16) (int32, error) {
	div := int32(asSignedWord(word1))
	if div == 0 {
		return 0, errors.New("divide by zero")
	}
	return int32(asSignedWord(word0)) % div, nil
}
