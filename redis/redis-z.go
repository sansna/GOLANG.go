package main

import (
	"fmt"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"

	rds "github.com/garyburd/redigo/redis"
)

func main() {
	cfg := cfgstruct.RedisSt{
		Hostport: "127.0.0.1:6379",
		Poolsize: 4,
		Timeout:  5,
	}
	cache, err := base.NewCacheV2(cfg, "redis")
	reply, err := cache.Zscore("1", "12")
	if err == rds.ErrNil {
		fmt.Print("err nil")
	} else {
		fmt.Println(reply, err)
	}

}
