package main
import "math/rand"
import "time"
import "fmt"
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}
