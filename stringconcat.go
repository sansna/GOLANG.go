package main
import "fmt"
import "strconv"
func main() {
	var a string
	a = strconv.FormatInt(19,10) +"," +strconv.FormatInt(192, 10)
	fmt.Println(a)
}
