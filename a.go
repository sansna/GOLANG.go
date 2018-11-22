package main
import "git.ixiaochuan.cn/xclib/common/utils"
import "time"
import "fmt"
func main() {
	fmt.Println(utils.TimestampToTimeString(time.Now().Unix()))
	fmt.Println(len("asdf"))
}
