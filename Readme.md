### Cute Print

a cute Go print package, simple, no dependencies!

### Installation
```zsh
go get -u github.com/zakaria-chahboun/cute
```
The idea is to simplify the printing of messages in the terminal, *exactly in case of error*. Letting you to put a title for your message in different colors.

### Print line
```go
	cute.Println("Hi everyone", "My name is Zakaria!")
```
![print line](./screenshots/01.png)

As you see the default color for the title is `yellow` and for the message is `purple`.

### Print muti-lines
```go
	cute.Println("Hi everyone", "My name is Zakaria!", "Zaki is my nick name.")
```
![print line](./screenshots/02.png)

### Check error
```go
    // equal to (if error != nil)
	cute.Check("Error", errors.New("This is a cute panic!"))
```
![print line](./screenshots/03.png)

### Change colors
You can change the colors:
* NoColor
* ColorRed 
* ColorYellow
* ColorGreen 
* ColorBlue  
* ColorPurple
* ColorCyan   
* ColorWhite 

```go
	cute.SetTitleColor(cute.ColorBlue)
	cute.SetMessageColor(cute.ColorGreen)
	cute.Println("Hi everyone", "My name is Zakaria!")
```
![print line](./screenshots/04.png)

### Printf
```go
	cute.Printf("Another title", "%s, a Country in North Africa.\n", "Morocco")
```
![print line](./screenshots/05.png)

-------
Twitter: [@Zaki_Chahboun](https://twitter.com/zaki_chahboun)
	