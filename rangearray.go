package main
import (
	"fmt"
)

func main() {
	a := []int64{1,2,3}
	// prints index of a
	for i := range a {
		fmt.Println(i)
	}
}
