package main
import (
	"fmt"
)

func main() {
	a := []int64{1,2,3}
	fmt.Println(a)
	b := []int64{5}
	a = append(b, a...)
	fmt.Println(a)
}
