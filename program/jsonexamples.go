package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {

	person := Person{Name: "vader", Age: 40, Email: "vader@instamojo.com"}
	marshelledData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(marshelledData))
}
