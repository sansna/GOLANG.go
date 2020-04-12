package main
import (
	"strings"
	"fmt"
)

func main() {
	strs := strings.SplitN("http://file.iupvideo.net/img/view/id/17612756", "/", 7)
	fmt.Println(len(strs), strs[6])
}
