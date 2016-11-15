package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("A random number %v\n", rand.Intn(10))
}
