package cute

import (
	"fmt"
	"os"
	"strings"
)

/* new type : color */
type color string

/* list of available colors */
const (
	NoColor = "\033[0m"

	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

/* local variables */
var (
	title_color   color = ColorYellow
	message_color color = ColorPurple
)

/* getter */
func SetTitleColor(c color) {
	title_color = c
}

/* gitter */
func SetMessageColor(c color) {
	message_color = c
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
func messageDraw(message string) (msg string) {
	msg = fmt.Sprintf("🭬 %s", message)
	return
}

/* println */
func Println(title string, messages ...string) {
	fmt.Printf("%v", title_color) // set the color
	fmt.Println(titleDraw(title))
	fmt.Printf("%v", message_color) // set the color
	for _, m := range messages {
		fmt.Println(messageDraw(m))
	}
	fmt.Printf("%v", NoColor) // set no color
}

/* println */
func Printf(title string, message string, params ...any) {
	fmt.Printf("%v", title_color) // set the color
	fmt.Println(titleDraw(title))
	fmt.Printf("%v", message_color) // set the color
	fmt.Printf(messageDraw(message), params...)
	fmt.Printf("%v", NoColor) // set no color
}

/* a cute panic like */
func Check(title string, err error) {
	if err != nil {
		// print title
		fmt.Printf("%v", ColorYellow)
		fmt.Println(titleDraw(title))

		// print error message
		fmt.Printf("%v", ColorRed)
		fmt.Println("🗲", err.Error())
		fmt.Printf("%v", NoColor)
		os.Exit(1)
	}
}
