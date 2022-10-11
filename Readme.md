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

Screenshot

<img src="./screenshots/01.png" alt="print line" width=500/>

> As you see the default color for the title is `yellow` and for the message is `purple`.

### Print muti-lines
```go
	cute.Println("Hi everyone", "My name is Zakaria!", "Zaki is my nick name.")
```

Screenshot

<img src="./screenshots/02.png" alt="print line" width=500/>

### Check error
```go
    // equal to (if error != nil)
	cute.Check("Error", errors.New("This is a cute panic!"))
```

Screenshot


<img src="./screenshots/03.png" alt="print line" width=500/>

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

Screenshot


<img src="./screenshots/04.png" alt="print line" width=500/>

### Printf
```go
	cute.Printf("Another title", "%s, a Country in North Africa.\n", "Morocco")
```

Screenshot


<img src="./screenshots/05.png" alt="print line" width=500/>

-------
Twitter: [@Zaki_Chahboun](https://twitter.com/zaki_chahboun)
	