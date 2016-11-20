package main

import "fmt"

func main() {
	// simple declaration
	var someSlice []int  // initialised with zero length and size
	fmt.Println(someSlice, len(someSlice))

	// initialisation with make using short hand
	someSlice2 := make([]int, 10)
	fmt.Println(someSlice2, len(someSlice2))

	// appending elements
	someSlice2 = append(someSlice2, 11)
	fmt.Println(someSlice2, len(someSlice2))
}
