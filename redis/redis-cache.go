package main

import (
	"fmt"

	rds "github.com/garyburd/redigo/redis"

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

	// get data if not exist returns nil
	buf, err := cache.GetString("asdfsdfsdf")
	if err == rds.ErrNil {
		fmt.Println("get null", buf, err)
	}

	// setexnx err always nil
	// succ if reply == "OK"
	reply, err := cache.SetExNx("asdfsdf", 1, 10)
	if err == rds.ErrNil {
		fmt.Println("setexnx:", reply, err)
	}
	if reply == "OK" {
		fmt.Println("ok")
	}

	// n == 1
	n, err := cache.Incrby("asdf", 1)
	fmt.Println(n, err)
	cache.Expire("asdf", 20)

	// cnt == 1
	cache.Set("a", "1")
	cnt, err := cache.DelRet("a")
	fmt.Println("delret:", cnt, err)

	// cnt == 2
	cache.Set("a", "1")
	cache.Set("b", "1")
	cnt, err = cache.MDelRet("a", "b", "B")
	fmt.Println("mdelret:", cnt, err)

	// 
	cache.Set("a", "1")
	ok, err := cache.Exists("a")
	fmt.Println("exists: ", ok, err)
	cache.Del("a")
}
