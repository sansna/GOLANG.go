package main
import (
	"fmt"
)
func main() {
	sdep := 0
	a := make(map[string]int)
	a = nil
	for _, cnt := range a {
		if sdep < cnt {
			sdep = cnt
		}
	}
	fmt.Println(sdep)
}
