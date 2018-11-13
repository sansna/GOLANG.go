package main
import "fmt"
func main() {
	a := fmt.Sprintf("%010d%s", 10239710, "life")
	fmt.Println(a)
	var digit int
	var str string
	fmt.Sscanf(a, "%010d%s", &digit, &str)
	fmt.Printf("%010d,%s\n", digit, str)
}
