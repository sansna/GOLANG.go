package main

import "fmt"
import "errors"

func newerr() (int, error) {
	return 1, errors.New("again!")
}

func cover() (mids []int64, s string, err error) {
	// the err is covered and returned.
	a, err := newerr()
	fmt.Println("asdf", a)
	return
}

func sss() (mids []int64, s string, err error) {
	//mids = []int64{}
	//s = string{}
	err = errors.New("wow")
	return
}

func main() {
	a, b, er := sss()
	fmt.Println(a, b, er)
	c, d, e := cover()
	fmt.Println(c, d, e)
}
