package main
import (
	"fmt"

	"git.ixiaochuan.cn/xclib/common/xctoken"
)

func main() {
	t, _ := xctoken.DecodeToken(
		"TbK9N-X7P-NdE-GjF9zAlHRy1DhIn7HO7mvm63X8BXf8rSPE=",
	)
	fmt.Println(t)
}
