## Cute Print

<img src="https://raw.githubusercontent.com/zakaria-chahboun/ZakiQtProjects/master/IMAGE1.png">

A cute Go print package *(fmt alternative)*, Minimalist, No dependencies!

Support me to be an independent open-source programmer 💟

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/U7U3FQ2JA)

## Installation
```sh
go get -u github.com/zakaria-chahboun/cute@latest
```
The idea is to simplify the printing of messages in the terminal, *especially in error case*. Letting you to put a title for your message in different colors.

## Print line
```go
cute.Println("Hi everyone", "My name is", "Zakaria!")
```

<img src="./screenshots/01.png" alt="print line" width=500/>

> As you see the default color for the title is `bright yellow` and for the message is `bright purple`.

## Note
If the results does not appears well. We suggest you to use fonts like `MesloLGS NF` or `Fira Code` in your terminal!

## Print muti-lines
```go
cute.Printlns("Hi everyone", "My name is Zakaria!", "Zaki is my nick name.")
```

<img src="./screenshots/02.png" alt="print multi-lines" width=500/>

## Check errors
The `cute.Check(..)` is useful in case of errors, especially if you have a lot of functions in you code. It will help you to print a clear error code.

```go
// equal to (if error != nil)
cute.Check("Error Title", errors.New("This is a cute panic!"))
```

<img src="./screenshots/03.png" alt="check errors" width=500/>

## Change colors
You can change the color:
* `ResetColor`
* `DefaultColor`


* `Black` / `BrightBlack`
* `Red` / `BrightRed`
* `Yellow` / `BrightYellow`
* `Green` / `BrightGreen`
* `Blue` / `BrightBlue`
* `Purple` / `BrightPurple`
* `Cyan` / `BrightCyan`
* `White` / `BrightWhite`

```go
cute.SetTitleColor(cute.BrightBlue)
cute.SetMessageColor(cute.BrightGreen)
cute.Println("Hi everyone", "My name is Zakaria!")
```

<img src="./screenshots/04.png" alt="colors" width=500/>

## Printf
```go
cute.Printf("Another title", "%s, a Country in North Africa.\n", "Morocco")
```

<img src="./screenshots/05.png" alt="printf" width=500/>

## List
You can print a list of lines dynamically! You can also specify the color for each line:

```go
// juice recipe 🧃
list := cute.NewList(cute.BrightBlue, "Yummy Juice!")
list.Add(cute.BrightGreen, "1 avocado 🥑")
list.Add(cute.BrightRed, "4 strawberry 🍓")
list.Addf(cute.White, "%d ML %s", 500, "milk 🥛")
list.Print()
```

<img src="./screenshots/06.png" alt="print list with colors" width=500/>

## How you can use it with Scan?
This is a little example:

```go
cute.Printf("How old are you?", "")

var age int
_, err := fmt.Scanln(&age)
cute.Check("Error scan", err)

cute.Println("Info", "Your age is:", age)
```
<img src="./screenshots/07.png" alt="use cute with Scan" width=500/>

## Real example
[Tarjem](https://github.com/zakaria-chahboun/tarjem) is used our cute package, Here are some examples:

<img src="./screenshots/real-example.png" alt="real example" />

## Contribute 💟
Feel free to contribute or propose a feature or share your idea with us!

-------
Twitter: [@Zaki_Chahboun](https://twitter.com/zaki_chahboun)
	