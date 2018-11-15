package main
import "time"
import "fmt"
func main() {
	a := time.NewTimer(time.Second*3)
	fmt.Println(time.Now().Unix())
	b := <-a.C
	fmt.Println(b.Unix())
	fmt.Println(time.Now().Unix())
	nanotime := time.Now().UnixNano()
	fmt.Println(nanotime)
	fmt.Printf("%020d\n", nanotime)

	var d string
	var i int64
	str := fmt.Sprintf("%025d%s", nanotime, "helloworld")
	fmt.Sscanf(str, "%025d%s", &i, &d)
	fmt.Println(i, d)
}