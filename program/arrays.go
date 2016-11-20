package main

import "fmt"

func main() {
	someArray := [3]string{"one", "two", "three"}
	fmt.Println(someArray)

	// invalid raises index error
	someArray[4] = "four"
}
