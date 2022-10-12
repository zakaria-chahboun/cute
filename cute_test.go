package cute

import (
	"errors"
	"testing"
)

// Test equality of CuteColor type.
func TestCuteColorType(t *testing.T) {
	redColor := CuteColor("\033[31m")
	if ColorRed != redColor {
		t.Fail()
	}
}

// Test show case.
func TestCuteFunctions(t *testing.T) {
	Println("Title", "This is just test", 123456)
	Printf("Title", "My ID is %v\n", 123456)

	SetTitleColor(ColorBrightBlue)
	SetMessageColor(ColorBrightGreen)
	Println("Hi everyone", "My name is Zakaria!")

	Check("Error", errors.New("This is a cute panic!"))
}
