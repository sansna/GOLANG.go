package main

import (
	"encoding/json"
	"fmt"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

type PostOffsetSt struct {
	Pid    int64 `json:"pid"`
	Offset int64 `json:"offset"`
}

func main() {
	cfg := cfgstruct.RedisSt{
		Hostport: "127.0.0.1:6379",
		Poolsize: 4,
		Timeout:  5,
	}
	cache, _ := base.NewCacheV2(cfg, "redis")
	a := make([]*PostOffsetSt, 0)
	reply, _ := cache.EvalSHA("c030467cad276bf950c7634ef327d3548a9386b4", 0)
	if _, ok := reply.([]byte); !ok {
		fmt.Println("error")
		return
	}
	json.Unmarshal(reply.([]byte), &a)
	for i, v := range a {
		fmt.Println(i, v.Pid, v.Offset)
	}
	return
}
