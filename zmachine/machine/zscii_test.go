// Test the low-level zscii functions.
package machine

import "testing"

func Test_zsciiWords(t *testing.T) {
	if w0, w1, w2, last := zsciiWords(0b00000100, 0b00100001); w0 != 1 && w1 != 1 && w2 != 1 && last != false {
		t.Errorf("1,1,1 = %x,%x,%x", w0, w1, w2)
	}
}

func Test_v1(t *testing.T) {
	// Excerpts from V1 Infocom files, for ensuring correct decoding.

}
