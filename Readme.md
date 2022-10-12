## Cute Print

<img src="https://raw.githubusercontent.com/zakaria-chahboun/ZakiQtProjects/master/IMAGE1.png">

a cute Go print package *(fmt alternative)*, minimalist, no dependencies!

## Installation
```zsh
go get -u github.com/zakaria-chahboun/cute
```
The idea is to simplify the printing of messages in the terminal, *especially in error case*. Letting you to put a title for your message in different colors.

## Print line
```go
cute.Println("Hi everyone", "My name is Zakaria!")
```

Screenshot

<img src="./screenshots/01.png" alt="print line" width=500/>

> As you see the default color for the title is `bright yellow` and for the message is `bright purple`.

## Print muti-lines
```go
cute.Println("Hi everyone", "My name is Zakaria!", "Zaki is my nick name.")
```

Screenshot

<img src="./screenshots/02.png" alt="print line" width=500/>

## Check errors
The `cute.Check(..)` is useful in case of errors, especially if you have a lot of functions in you code. It will help you to print a clear error code.

```go
// equal to (if error != nil)
cute.Check("Error Title", errors.New("This is a cute panic!"))
```

Screenshot


<img src="./screenshots/03.png" alt="print line" width=500/>

## Change colors
You can change the color:
* `ColorReset`
* `ColorDefault`


* `ColorBlack` / `ColorBrightBlack`
* `ColorRed` / `ColorBrightRed`
* `ColorYellow` / `ColorBrightYellow`
* `ColorGreen` / `ColorBrightGreen`
* `ColorBlue` / `ColorBrightBlue`
* `ColorPurple` / `ColorBrightPurple`
* `ColorCyan` / `ColorBrightCyan`
* `ColorWhite` / `ColorBrightWhite`

```go
cute.SetTitleColor(cute.ColorBrightBlue)
cute.SetMessageColor(cute.ColorBrightGreen)
cute.Println("Hi everyone", "My name is Zakaria!")
```

Screenshot


<img src="./screenshots/04.png" alt="print line" width=500/>

## Printf
```go
cute.Printf("Another title", "%s, a Country in North Africa.\n", "Morocco")
```

Screenshot


<img src="./screenshots/05.png" alt="print line" width=500/>

-------
Twitter: [@Zaki_Chahboun](https://twitter.com/zaki_chahboun)
	