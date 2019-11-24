package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	MAX := 30000000
	N := 1000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		fmt.Println(rand.Intn(MAX))
	}
}
