package main
import (
	"sort"
	"fmt"
)

func main() {
	z := []int64{3,5,1,2,6,9,8,7}
	sort.Slice(z, func(i, j int) bool {
		return z[i] < z[j]
	})
	fmt.Println(z)
}
