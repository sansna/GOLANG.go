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

	reply, err := cache.SetExNx("asdf", 1, 1000)
	fmt.Println(reply.(string), err, reply == "OK")

	reply, err = cache.SetExNx("asdf", 1, 1000)
	fmt.Println(reply, err)
	cache.Del("asdf")
}
