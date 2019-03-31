package main
import "math/rand"
import "time"
import "fmt"
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Perm(1))
}
