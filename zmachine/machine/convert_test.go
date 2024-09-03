// Tests for the convert file.
package machine

import "testing"

func Test_asWord(t *testing.T) {
	if v := asWord(0, 0); v != 0 {
		t.Errorf("0000 = %x", v)
	}
	if v := asWord(0x10, 0xf0); v != 0x10f0 {
		t.Errorf("10f0 = %x", v)
	}
	if v := asWord(0xf0, 0x10); v != 0xf010 {
		t.Errorf("f010 = %x", v)
	}
}

func Test_fromWord(t *testing.T) {
	if v0, v1 := fromWord(0x0000); v0 != 0x00 && v1 != 0x00 {
		t.Errorf("0000 = %x,%x", v0, v1)
	}
	if v0, v1 := fromWord(0xf023); v0 != 0xf0 && v1 != 0x23 {
		t.Errorf("f023 = %x,%x", v0, v1)
	}
	if v0, v1 := fromWord(0x23f0); v0 != 0x23 && v1 != 0xf0 {
		t.Errorf("23f0 = %x,%x", v0, v1)
	}
}

func Test_asLong(t *testing.T) {
	if v := asLong(0, 0, 0, 0); v != 0 {
		t.Errorf("00000000 = %x", v)
	}
	if v := asLong(0x10, 0xf0, 0x23, 0xa0); v != 0x10f023a0 {
		t.Errorf("10f023a0 = %x", v)
	}
	if v := asLong(0xf0, 0x10, 0x23, 0xa0); v != 0xf01023a0 {
		t.Errorf("f01023a0 = %x", v)
	}
}

func Test_fromLong(t *testing.T) {
	if v0, v1, v2, v3 := fromLong(0x00000000); v0 != 0x00 && v1 != 0x00 && v2 != 0x00 && v3 != 0x00 {
		t.Errorf("00000000 = %x,%x,%x,%x", v0, v1, v2, v3)
	}
	if v0, v1, v2, v3 := fromLong(0x10f023a0); v0 != 0x10 && v1 != 0xf0 && v2 != 0x23 && v3 != 0xa0 {
		t.Errorf("10f023a0 = %x,%x,%x,%x", v0, v1, v2, v3)
	}
	if v0, v1, v2, v3 := fromLong(0xf01023a0); v0 != 0xf0 && v1 != 0x10 && v2 != 0x23 && v3 != 0xa0 {
		t.Errorf("f01023a0 = %x,%x,%x,%x", v0, v1, v2, v3)
	}
}

func Test_asSignedWord(t *testing.T) {
	if v := asSignedWord(0); v != 0 {
		t.Errorf("0 = %d", v)
	}
	if v := asSignedWord(0x7fff); v != 0x7fff {
		t.Errorf("7fff = %d", v)
	}
	if v := asSignedWord(0xffff); v != -1 {
		t.Errorf("ffff = %d", v)
	}
	if v := asSignedWord(0x8000); v != -32768 {
		t.Errorf("8000 = %d", v)
	}
}

func Test_normalizeSignedWord(t *testing.T) {
	if v := normalizeSignedWord(0); v != 0 {
		t.Errorf("0 = %d", v)
	}
	if v := normalizeSignedWord(0x7fff); v != 0x7fff {
		t.Errorf("7fff = %d", v)
	}
	if v := normalizeSignedWord(-1); v != 0xffff {
		t.Errorf("ffff = %d", v)
	}
	if v := normalizeSignedWord(-32768); v != 0x8000 {
		t.Errorf("8000 = %d", v)
	}
	if v := normalizeSignedWord(-10); v != 0xfff6 {
		t.Errorf("fff6 = %x", v)
	}
	if v := normalizeSignedWord(-2); v != 0xfffe {
		t.Errorf("fffe = %x", v)
	}
	if v := normalizeSignedWord(-13); v != 0xfff3 {
		t.Errorf("fff3 = %x", v)
	}
	if v := normalizeSignedWord(-5); v != 0xfffb {
		t.Errorf("fffb = %x", v)
	}
}
