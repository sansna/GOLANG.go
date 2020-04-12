package main

import (
	"fmt"
	"git.ixiaochuan.cn/gateway/controllers"
)

func main() {
	s := "凑凑字数凑字数凑字数凑字数凑字数凑字数字数"
	out := controllers.DeDuplicateString(s)
	fmt.Println(out)
}
