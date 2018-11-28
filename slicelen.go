package main

import "fmt"

func main() {
	a := make([]int64,0,10)
	b := make([]int64,10)
	a = append(a,1)
	fmt.Println(len(a),len(b))
}
