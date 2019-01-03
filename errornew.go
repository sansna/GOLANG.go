package main
import "errors"
import "fmt"
func sss() (err error) {
	err = errors.New("Connection Error")
	return
}
func main() {
	err := sss()
	if err.Error() == "Connection Error" {
		fmt.Println("done")
	}
	if err.Error() != "Connection Error" {
		fmt.Println("ok")
	}
}
