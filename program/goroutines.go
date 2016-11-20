package main

import (
	"fmt"
	"time"
)

func StartWork(index int) {
	for {
		fmt.Println("Worker", index)
	}

}

func main() {
	for i := 0; i < 3; i++ {
		go StartWork(i)
	}

	time.Sleep(5 * time.Second)
}
