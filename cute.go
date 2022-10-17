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
	ResetColor   CuteColor = "\033[0m"
	DefaultColor CuteColor = "\033[39m"

	Black  CuteColor = "\033[30m"
	Red    CuteColor = "\033[31m"
	Green  CuteColor = "\033[32m"
	Yellow CuteColor = "\033[33m"
	Blue   CuteColor = "\033[34m"
	Purple CuteColor = "\033[35m"
	Cyan   CuteColor = "\033[36m"
	White  CuteColor = "\033[37m"

	BrightBlack  CuteColor = "\033[90m"
	BrightRed    CuteColor = "\033[91m"
	BrightGreen  CuteColor = "\033[92m"
	BrightYellow CuteColor = "\033[93m"
	BrightBlue   CuteColor = "\033[94m"
	BrightPurple CuteColor = "\033[95m"
	BrightCyan   CuteColor = "\033[96m"
	BrightWhite  CuteColor = "\033[97m"
)

/* local variables */
var (
	current_title_color   CuteColor = BrightYellow
	current_message_color CuteColor = BrightPurple
)

/* Change the title color */
func SetTitleColor(c CuteColor) {
	current_title_color = c
}

/* Change the message color */
func SetMessageColor(c CuteColor) {
	current_message_color = c
}

/* Printlns multi-lines */
func Printlns(title string, messages ...any) {
	// set the color
	colorize(current_title_color)
	// print title
	fmt.Println(drawTitle(title))
	// set the color
	colorize(current_message_color)
	// print messages
	for _, m := range messages {
		fmt.Println(drawMessage(m))
	}
	// reset
	colorize(ResetColor)
}

/* Println */
func Println(title string, messages ...any) {
	// set the color
	colorize(current_title_color)
	// print title
	fmt.Println(drawTitle(title))
	// set the color
	colorize(current_message_color)
	// print messages
	if len(messages) != 0 {
		lines := fmt.Sprintln(messages...)
		fmt.Println(drawMessage(lines))
	}
	// reset
	colorize(ResetColor)
}

/* Println */
func Printf(title string, message string, params ...any) {
	// set the color
	colorize(current_title_color)
	// print title
	fmt.Println(drawTitle(title))
	// set the color
	colorize(current_message_color)
	// print messages
	fmt.Printf(drawMessage(message), params...)
	// reset
	colorize(ResetColor)
}

/* a cute panic like */
func Check(title string, err error) {
	if err != nil {
		// set red color for all
		fmt.Printf("%v", BrightRed)
		// print title
		fmt.Println(drawTitle(title))
		// print error message
		fmt.Println("🗲", err.Error())
		colorize(ResetColor)
		os.Exit(1)
	}
}

/* local draw title with box */
func drawTitle(title string) (box string) {
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
func drawMessage(message any) (msg string) {
	msg = fmt.Sprintf("🭬 %v", message)
	return
}

/* local: start coloring */
func colorize(color CuteColor) {
	fmt.Printf("%v", color)
}
