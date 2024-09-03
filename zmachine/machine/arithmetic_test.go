// Test the arithmetic operations.
package machine_test

import (
	"testing"

	"github.com/groboclown/go-zmachine/zmachine/machine"
)

func Test_signedDiv(t *testing.T) {
	if v, _ := machine.ArithmeticDivide(0xfff6, 2); v != -5 {
		t.Errorf("-11 / 2 = %v", v)
	}
	if v, _ := machine.ArithmeticDivide(0xfff6, 0xfffe); v != 5 {
		t.Errorf("-11 / -2 = %v", v)
	}
	if v, _ := machine.ArithmeticDivide(11, 0xfffe); v != -5 {
		t.Errorf("11 / -2 = %v", v)
	}

	if v, _ := machine.ArithmeticRemainder(0xfff3, 5); v != -3 {
		t.Errorf("-13 / 5 = %v", v)
	}
	if v, _ := machine.ArithmeticRemainder(0xfff3, 0xfffb); v != -3 {
		t.Errorf("-13 / -5 = %v", v)
	}
	if v, _ := machine.ArithmeticRemainder(13, 0xfffb); v != 3 {
		t.Errorf("13 / -5 = %v", v)
	}
}
