package main

import (
	"fmt"
	"strconv"

	"git.ixiaochuan.cn/xclib/common/freqlimit"
	"git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

func main() {
	// init
	cfg := cfgstruct.RedisSt{
		Hostport: "127.0.0.1:7001",
		Poolsize: 4,
		Timeout:  5,
	}
	var rl [3]*freqlimit.RateLevel
	rl[0] = &freqlimit.RateLevel{
		Base:     7,
		Max:      15,
		Dur:      60,
		Name:     "min",
		Rec1:     1,
		Rec1Name: "rpl",
		Rec2:     20,
		Rec2Name: "old",
	}
	rl[1] = &freqlimit.RateLevel{
		Base:     20,
		Max:      100,
		Dur:      3600,
		Name:     "hour",
		Rec1:     1,
		Rec1Name: "rpl",
		Rec2:     20,
		Rec2Name: "old",
	}
	rl[2] = &freqlimit.RateLevel{
		Base:     40,
		Max:      300,
		Dur:      86400,
		Name:     "day",
		Rec1:     1,
		Rec1Name: "rpl",
		Rec2:     20,
		Rec2Name: "old",
	}
	param := &freqlimit.FreqLimitParam{
		DoSHAKey:   "728ad993086c57a8c8717cbe345f1803c1c2c2ac",
		Prefix:     "cs",
		RL:         rl,
		BanKey:     "ban",
		DefaultRet: int64(0),
	}
	l, err := freqlimit.Init(&cfg, param)
	if err != nil {
		fmt.Println("error: %v", err)
		return
	}
	// do
	midStr := strconv.FormatInt(123, 10)
	keys := l.GenerateKeys(midStr)
	fmt.Println(keys)
	for i := 0; i < 10; i++ {
		fmt.Println(i, l.Do(midStr))
	}
	fmt.Println(l.IncAllLevelRecKey(midStr, 1))
	return
}
