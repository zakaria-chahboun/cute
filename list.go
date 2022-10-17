package cute

import "fmt"

/* CuteList (print lines dynamically) */
type CuteList struct {
	Title      string
	TitleColor CuteColor
	Lines      []CuteLine
}

/* CluteLine (child of CuteList) */
type CuteLine struct {
	Message      any
	MessageColor CuteColor
}

// constructor of list
func NewList(color CuteColor, title string) *CuteList {
	return &CuteList{
		TitleColor: color,
		Title:      title,
	}
}

/* add message to list */
func (this *CuteList) Add(color CuteColor, message any) {
	this.Lines = append(this.Lines, CuteLine{
		MessageColor: color,
		Message:      message,
	})
}

/* addf message to list */
func (this *CuteList) Addf(color CuteColor, message string, params ...any) {
	this.Lines = append(this.Lines, CuteLine{
		MessageColor: color,
		Message:      fmt.Sprintf(message, params...),
	})
}

/* print list */
func (this *CuteList) Print() {
	// set title color
	colorize(this.TitleColor)
	// print title
	fmt.Println(drawTitle(this.Title))

	for _, line := range this.Lines {
		// set message color
		colorize(line.MessageColor)
		// print message
		fmt.Println(drawMessage(line.Message))
	}
	// reset
	colorize(ResetColor)
}
