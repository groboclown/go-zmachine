// Test the instructions parsing
package machine_test

import (
	"testing"

	"github.com/groboclown/go-zmachine/zmachine/machine"
)

func Test_DecodeVarType(t *testing.T) {
	if v := machine.DecodeVarType(0b00101111); v[0] != machine.LargeOperand && v[1] != machine.VariableOperand && v[2] != machine.OmittedOperand && v[3] != machine.OmittedOperand {
		t.Errorf("00,10,11,11 = %x,%x,%x,%x", v[0], v[1], v[2], v[3])
	}
}
