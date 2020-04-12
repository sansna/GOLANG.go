package main
import (
	"math"
	"fmt"
)
func main() {
	factor := 1 / (math.Ln2 * math.Log2(math.E+3600))
	fmt.Println(factor*0.3)
	fmt.Println(math.Log2(1.1+float64(2)/10))
}
