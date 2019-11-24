package main

import (
	"fmt"

	"git.ixiaochuan.cn/xclib/common/xccrypt"
	"git.ixiaochuan.cn/xclib/common/xckey"
)

func main() {
	auth := "crypt-124e9177ac8b466ee3c8c6713cf7a8f06704ed13952238941ba9ecd1db83f840"
	fmt.Println(xccrypt.Decrypt(auth, xckey.AppKey))
}
