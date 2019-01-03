package main

import "fmt"
import "errors"

func aa() (err error) {
	a, err := ss()
	fmt.Printf("a is %d\n", a)
	return
}

func ss() (a int, err error) {
	return 1, errors.New("wow")
}

func main() {
	err := aa()
	fmt.Println(err)
}
