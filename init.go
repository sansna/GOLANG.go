package main
import "bytes"
import "fmt"
func main() {
	for {
		// inits a to nil
		var a bytes.Buffer
		a.WriteString("s")
		fmt.Println(a.String())
	}
}
