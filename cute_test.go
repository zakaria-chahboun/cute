package cute

import "testing"

// Test equality of CuteColor type.
func TestCuteColorType(t *testing.T) {
	Println("Title", "This is just test", 123456)
	SetTitleColor(ColorBrightBlue)
	SetMessageColor(ColorBrightGreen)
	Println("Hi everyone", "My name is Zakaria!")
	redColor := CuteColor("\033[31m")
	if ColorRed != redColor {
		t.Fail()
	}
}
