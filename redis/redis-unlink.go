package main

import (
	"fmt"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

func main() {
	cfg := cfgstruct.RedisSt{
		Hostport: "127.0.0.1:6379",
		Poolsize: 4,
		Timeout:  5,
	}
	cache, err := base.NewCacheV2(cfg, "redis")
	cache.SetPrefix("sdfs_")
	cache.Set("1", "100")
	cache.Set("2", "100")
	cache.Set("3", "100")
	cache.Unlink("1","2","3")
	val, err := cache.GetData("1")
	fmt.Println(string(val), err)
}
