package main
import "fmt"
func main() {
	a := make(map[string]interface{})
	a["lo"] = ""
	fmt.Println(a)

	i := make(map[int64]int64)
	fmt.Println(i[10])
}
