package cute

import (
	"fmt"
	"os"
	"strings"
)

/* new type : color */
type CuteColor string

/* list of available colors */
const (
	ColorReset   CuteColor = "\033[0m"
	ColorDefault CuteColor = "\033[39m"

	ColorBlack  CuteColor = "\033[30m"
	ColorRed    CuteColor = "\033[31m"
	ColorGreen  CuteColor = "\033[32m"
	ColorYellow CuteColor = "\033[33m"
	ColorBlue   CuteColor = "\033[34m"
	ColorPurple CuteColor = "\033[35m"
	ColorCyan   CuteColor = "\033[36m"
	ColorWhite  CuteColor = "\033[37m"

	ColorBrightBlack  CuteColor = "\033[90m"
	ColorBrightRed    CuteColor = "\033[91m"
	ColorBrightGreen  CuteColor = "\033[92m"
	ColorBrightYellow CuteColor = "\033[93m"
	ColorBrightBlue   CuteColor = "\033[94m"
	ColorBrightPurple CuteColor = "\033[95m"
	ColorBrightCyan   CuteColor = "\033[96m"
	ColorBrightWhite  CuteColor = "\033[97m"
)

/* local variables */
var (
	current_title_color   CuteColor = ColorBrightYellow
	current_message_color CuteColor = ColorBrightPurple
)

/* Change the title color */
func SetTitleColor(c CuteColor) {
	current_title_color = c
}

/* Change the message color */
func SetMessageColor(c CuteColor) {
	current_message_color = c
}

/* Println */
func Println(title string, messages ...interface{}) {
	fmt.Printf("%v", current_title_color) // set the color
	fmt.Println(titleDraw(title))
	fmt.Printf("%v", current_message_color) // set the color
	for _, m := range messages {
		fmt.Println(messageDraw(m))
	}
	fmt.Printf("%v", ColorReset) // set no color
}

/* Println */
func Printf(title string, message string, params ...interface{}) {
	fmt.Printf("%v", current_title_color) // set the color
	fmt.Println(titleDraw(title))
	fmt.Printf("%v", current_message_color) // set the color
	fmt.Printf(messageDraw(message), params...)
	fmt.Printf("%v", ColorReset) // set no color
}

/* a cute panic like */
func Check(title string, err error) {
	if err != nil {
		// print title
		fmt.Printf("%v", ColorBrightYellow)
		fmt.Println(titleDraw(title))

		// print error message
		fmt.Printf("%v", ColorBrightRed)
		fmt.Println("🗲", err.Error())
		fmt.Printf("%v", ColorReset)
		os.Exit(1)
	}
}

/* local drawing title */
func titleDraw(title string) (box string) {
	// box elements
	topright := "╮"
	topleft := "╭"
	w := "─"
	h := "│"
	bottomright := "╯"
	bottomleft := "╰"

	// draw line
	line := strings.Repeat(w, len(title)+2)

	// print title box
	box = fmt.Sprintf("%v%v%v\n", topleft, line, topright)
	box += fmt.Sprintf("%v %v %v\n", h, title, h)
	box += fmt.Sprintf("%v%v%v", bottomleft, line, bottomright)
	return
}

/* local draw message */
func messageDraw(message interface{}) (msg string) {
	msg = fmt.Sprintf("🭬 %v", message)
	return
}
