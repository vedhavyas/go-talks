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
	data := `{"name":"vader","age":300,"email":"luke@instamojo.com"}`
	person := Person{}
	err := json.Unmarshal([]byte(data), &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", person)
}
