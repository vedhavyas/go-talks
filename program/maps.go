package main

import "fmt"

func main() {
	someMap := make(map[string]int)  // initialisation

	someMap["some_key"] = 12 // insert

	value := someMap["some_key"] // retrieve
	fmt.Println(value)

	value, ok := someMap["some_key"] // conditional retrieve
	fmt.Println(ok, value)

	value, ok = someMap["unknown"] // conditional retrieve
	fmt.Println(ok, value)

	delete(someMap, "some_key") // delete
	fmt.Println(someMap)
}
