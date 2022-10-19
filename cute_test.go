package cute

import (
	"errors"
	"testing"
)

// Test equality of CuteColor type.
func TestCuteColorType(t *testing.T) {
	redColor := CuteColor("\033[31m")
	if Red != redColor {
		t.Fail()
	}
}

// Test show case.
func TestCuteFunctions(t *testing.T) {
	Println("Println", "Hi! i'm", "Zaki")
	Printf("Printf", "UID is %v\n", 1399745)

	SetTitleColor(BrightBlue)
	SetMessageColor(BrightGreen)
	Printlns("Printlns", "Line1", "Line2", "Line3")

	Println("Println: No lines")
	Printlns("Printlns: No lines")

	list := NewList(BrightBlue, "Yummy Juice!")
	list.Add(BrightRed, "4 strawberry")
	list.Add(BrightGreen, "1 avocado")
	list.Addf(White, "%d ml %s", 500, "milk")
	list.Print()

	Check("Error", errors.New("This is a cute panic!"))
}

// test unicode in title
func TestCuteTitleUnicode(t *testing.T) {
	// arabic
	Println("﴾ الله خالِقُ كُلِّ شيء ﴿")
	// frensh
	Println("délicieux pain français")
	// spanish
	Println("¡Hola! español")
	// german
	Println("ich möchte Kaffee")
	// chinese (simplified)
	Println("中国的长城")
	// chinese (traditional)
	Println("中國的長城")
	// japanese (hiragana)
	Println("進撃の巨人 すばらしい")
	// russian
	Println("русская литература")
	// turkish
	Println("türk şarküteri")
	// korean
	Println("한국라면 맛있다")
	// hindi (devanagari)
	Println("मसालों की भूमि")
	// emoji
	Println("✅  🚀")

	// complex
	Println("sweet | sucré | süß | 甜的 | 甘い | मीठा | حلو | 🍬")
}
