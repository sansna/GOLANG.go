package main
import "time"
import "fmt"
func main() {
	t := time.Unix(0,0)
	fmt.Println(int(time.Since(t).Months())/24/365)
}
