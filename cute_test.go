package cute

import "testing"

// Test equality of CuteColor type.
func TestCuteColorType(t *testing.T) {
	redColor := CuteColor("\033[31m")
	if ColorRed != redColor {
		t.Fail()
	}
}