package main

import (
	"fmt"
)

// iota starts from 0
const (
	A = iota
	B
	C
	D
)

const (
	AA = 1 << iota
	BB
	CC
	DD
	EE
)

func main() {
	fmt.Println(A, B, C)
	fmt.Println(AA, BB, CC)
}
