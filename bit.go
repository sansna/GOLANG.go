package main
import "math"
import "fmt"
func main() {
	for i:=0; i < 64; i++ {
		fmt.Printf("%064b\n", uint64(math.Pow(2,float64(i))))
	}
	var a int64
	a = math.MaxInt64
	fmt.Printf("%064b\n", a)

	fmt.Println(1<<4)
}
