package main

import (
	"math"
	"time"
	"fmt"
)

func main() {
	st := time.Now().Nanosecond()
	a := math.Exp(1)
	et := time.Now().Nanosecond()
	fmt.Print(et-st, a)
}
