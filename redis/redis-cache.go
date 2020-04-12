package main

import (
	"fmt"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/xcproto"
	"git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

func main() {
	cfg := cfgstruct.RedisSt{
		Hostport: "127.0.0.1:6379",
		Poolsize: 4,
		Timeout:  5,
	}
	cache, _ := base.NewCacheV2(cfg, "redis")
	res, _ := cache.GetData("cache_post_base_post_2395344")
	post := new(xcproto.PostSt)
	post.Unmarshal(res)
	fmt.Println(post)
}
