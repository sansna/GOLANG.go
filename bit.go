package main
import "math"
import "fmt"
func main() {
	for i:=0; i < 64; i++ {
		fmt.Printf("%064b\n", uint64(math.Pow(2,float64(i))))
	}
}
