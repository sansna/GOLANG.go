package main
import (
	"fmt"
)
func main() {
	a := []int64{1,2,3}
	a[2]--
	a[0]--
	fmt.Println(a)
}
