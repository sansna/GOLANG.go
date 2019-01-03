package main

import "fmt"

func sss(next int64) (ret int64, err error) {
	fmt.Println("inside", next)
	return next + 1, nil
}
func main() {
	next := int64(1)
	var err error
	for {
		next, err = sss(next)
		fmt.Println("err:", err)
		fmt.Println("outside", next)
		if next == 5 {
			break
		}
	}
}
