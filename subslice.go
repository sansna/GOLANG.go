package main
import (
	"fmt"
)

func main() {
	a := []int64{1,2,3,4,5,6,7,8}
	fmt.Println(a[0:3])
	// This outbounds slice size.
	//fmt.Println(a[0:30])
	fmt.Println(a[:3])
	fmt.Println(a[:])
	fmt.Println(a[1:])
}
