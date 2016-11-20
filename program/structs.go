package main

import "fmt"

type Rectangle struct {
	Length  int
	Breadth int
}

func main() {
	rectangle := Rectangle{Length: 10, Breadth: 20}
	fmt.Println(rectangle.Breadth)
	fmt.Println(rectangle.Length)
	fmt.Println(rectangle)
}
