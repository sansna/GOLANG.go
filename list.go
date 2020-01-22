package main
import (
	"fmt"
)

var b [2]map[string]interface{}

func main() {
	a := []int64{1,2,3}
	a[2]--
	a[0]--
	fmt.Println(a)

	//b = append(b, make(map[string]interface{}))
	fmt.Println(b[0], len(b))
}
