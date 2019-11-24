package main
import (
	"fmt"
)

func Genmap() map[int64]int64 {
	a := map[int64]int64 {
		1:2,
		3:4,
	}
	return a
}
func main() {
	z:=Genmap()
	// takes effect
	z[3]=6
	z[7]=5

	fmt.Println(z)
}
