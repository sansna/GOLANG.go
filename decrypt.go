package main

import (
	"fmt"

	//"git.ixiaochuan.cn/xclib/common/xccrypt"
	//"git.ixiaochuan.cn/xclib/common/xckey"
	"git.ixiaochuan.cn/xc_server/xclib/common/xccrypt"
	"git.ixiaochuan.cn/xc_server/xclib/common/xckey"
)

func main() {
	auth := "crypt-6c8bff0384ac97d9e9e526d46966fec3dc3acb74a00a693c7cac626391de0f79"
	fmt.Println(xccrypt.Decrypt(auth, xckey.RedisKey))
	
	dest := "b7sameEdxrfDiu3a"
	fmt.Println(xccrypt.Encrypt(dest, xckey.RedisKey))
}
