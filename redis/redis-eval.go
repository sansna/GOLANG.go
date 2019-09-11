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
	reply, err := cache.EvalSHA("3338e356cd10308d0e5e239e5b0ec9bf95ac417d", 10, "min", "hour", "day", 1, 2, 3, 4, 5, 6, "user", 60, 3600, 86400, 20, 20, 20, 1, 1, 1, 7, 20, 40, 15, 100, 300)
	//reply, err := cache.EvalSHA("3b3974240902b9633f561abfd5f384fa29339963", 3, "min", "hour", "day", 30, 120, 3600)
	fmt.Println(reply.(int64), err)
}
