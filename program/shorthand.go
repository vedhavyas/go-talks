package main

import "fmt"

func main() {
	// short hand with :=
	// the type is inferred from the value asigned
	shortHand := true
	fmt.Println(shortHand)

	// zero values of common types
	var a int
	var b string
	var c float32
	var d bool
	fmt.Println("Int - ", a) // will be equal to 0
	fmt.Println("String - ", b) // will be equal to ""
	fmt.Println("Float32 - ", c) // will be equal to 0
	fmt.Println("Boolean - ", d) // will be equal to false
}
