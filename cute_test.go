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
	Println("Println", "Hi! i'm", "Zaki")
	Printf("Printf", "UID is %v\n", 1399745)

	SetTitleColor(ColorBrightBlue)
	SetMessageColor(ColorBrightGreen)
	Printlns("Printlns", "Line1", "Line2", "Line3")

	Println("Println: No lines")
	Printlns("Printlns: No lines")

	list := NewList(ColorBrightBlue, "Yummy Juice!")
	list.Add(ColorBrightRed, "4 strawberry")
	list.Add(ColorBrightGreen, "1 avocado")
	list.Addf(ColorWhite, "%d ml %s", 500, "milk")
	list.Print()

	Check("Error", errors.New("This is a cute panic!"))
}
