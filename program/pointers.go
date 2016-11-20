package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) SetNamePointer(name string) {
	person.Name = name
}

func (person Person) SetName(name string) {
	person.Name = name
}

func main() {
	darth := Person{Name: "Darth vader"}
	darth.SetName("Anakin")
	fmt.Println(darth.Name)

	luke := &Person{Name: "Luke Skywalker"} //luke is pointer reference
	luke.SetNamePointer("Leia")
	fmt.Println(luke.Name)
}
