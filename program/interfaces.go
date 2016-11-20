package main

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	Side int
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func printArea(shape Shape) {
	fmt.Println("Area is -", shape.Area())
}

func main() {
	square := Square{Side: 5}
	printArea(square)
}
