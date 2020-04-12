package main
import "strings"
import "fmt"
func main() {
	a := "s,,3,4,"
	strs:=strings.Split(a, ",")
	fmt.Println(len(strs))
	for i:=0; i < len(strs);i ++ {
		fmt.Println(strs[i])
	}

	url := "http://file.iupvideo.net/img/view/id/30377103"
	list := strings.Split(url, "/")
	fmt.Println(list[6])
}
