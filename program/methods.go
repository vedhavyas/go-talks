package main

import "fmt"

type Rectangle struct {
	Length  int
	Breadth int
}

func (rect Rectangle) Area() int {
	return rect.Breadth * rect.Length
}

func main() {
	rectangle := Rectangle{Length: 10, Breadth: 20}
	fmt.Println("Area is -", rectangle.Area())
}
