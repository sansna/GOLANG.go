package main
import (
	"fmt"
)

func main() {
	help_dedup := make(map[int64]bool)
	a := []int64{1,2,3,45,1}
	r := []int64{}
	for _, i := range a {
		if !help_dedup[i] {
			r = append(r, i)
			help_dedup[i] = true
		}
	}
	fmt.Print(r)
}
