package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	testmap := make(map[string]string)

	testmap["haha"] = "hello"

	data, err := json.Marshal(testmap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	Test(1, "a", "b", "c")
}

func Test(test1 int, test ...[]string) {

}
