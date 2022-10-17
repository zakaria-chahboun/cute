package cute

import "fmt"

/* List (print lines dynamically) */
type List struct {
	title string
	lines []map[CuteColor]string // color, message
}

// constructor of list
func NewList(title string) *List {
	return &List{
		title: title,
		lines: []map[CuteColor]string{},
	}
}

/* add message to list*/
func (this *List) Add(message any) {
	line := map[CuteColor]string{
		current_message_color: messageDraw(fmt.Sprint(message)),
	}
	this.lines = append(this.lines, line)
}

/* addf message to list */
func (this *List) Addf(message string, params ...any) {
	line := map[CuteColor]string{
		current_message_color: fmt.Sprintf(messageDraw(message), params...),
	}
	this.lines = append(this.lines, line)
}

/* print list of messages lines */
func (this *List) Print() {
	// set color
	colorize(current_title_color)
	// print title
	fmt.Println(titleWithBoxDraw(this.title))
	for _, line := range this.lines {
		for color, message := range line {
			// set color
			colorize(color)
			// print message
			fmt.Println(message)
		}
	}
	// reset
	colorize(ColorReset)
}
