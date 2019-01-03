package main
import "fmt"
func sss(a string, b ...int64) {
	args := []interface{}{}
	for _, v := range b {
		args = append(args, v)
	}
	fmt.Println(a, args)
	return
}
func main() {
	sss("li",1,2,3,4,5)
	sss("li")
	return
}
