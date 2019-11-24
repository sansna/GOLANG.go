package main

import (
	"fmt"
)

func main() {
	v := []string{"sl", "zxc", "12j"}
	for _, s := range v {
		fmt.Println(s)
	}
	for _, i := range []int64{1,2,3,4,5} {
		fmt.Println(i)
	}
}
