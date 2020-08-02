package main

import (
	"encoding/json"
	"fmt"
	"git.ixiaochuan.cn/xclib/common/logger"
	"git.ixiaochuan.cn/xclib/common/xcabtest"
	"git.ixiaochuan.cn/xclib/common/xcabtest/cfgstruct"
	"github.com/mitchellh/mapstructure"
	//"strconv"
	//"git.ixiaochuan.cn/pp_server/live/gateway/controller"
)

type GiftTypeCount struct {
	Typeid string
	Count  int64
}
type PromptConfig struct {
	At         bool  // 1.本处是否下发
	Intval     int64 // 2.本处的间隔 (0: 无间隔，-1无穷间隔，>0：间隔秒)
	DirectSend int64 // 3.本处满足间隔时间后礼物下发方式(如果1，就不用设置key存放礼包信息直接下发,如果-1:不用下发且不用存放礼包信息仅提示礼包信息,0:正常存放礼包信息)
}
type GiftConfig struct {
	Config      []GiftTypeCount // ab礼包物品情况
	Count       int64           // 领取本礼包次数上限 (-1: 无穷领取次数，0:不能领取，1:最多一次...)
	PackName    string          // 本礼包名称
	RoomConfig  PromptConfig    // 房间配置
	IndexConfig PromptConfig    // 首页配置
}

func GetGiftPackParam(place string) (ret map[string]GiftConfig) {
	fun := "GetGiftPackParam"
	// map packname -> config
	ret = make(map[string]GiftConfig)
	out := make(map[string]GiftConfig)
	mapinfo := xcabtest.GetMap("default", "live_giftpack_config", "info", nil)
	if mapinfo != nil {
		err := mapstructure.Decode(mapinfo, &out)
		if err != nil {
			logger.Error("%s: fail get dec: %v", fun, mapinfo)
		}
	}

	for giftpackname, conf := range out {
		switch place {
		case "room":
			if !conf.RoomConfig.At {
				continue
			}
		case "index":
			if !conf.IndexConfig.At {
				continue
			}
		default: // all
		}
		conf.PackName = giftpackname
		ret[giftpackname] = conf
	}

	logger.Info("%s: got conf: %v", fun, out)

	return
}
func main() {
	cfg := cfgstruct.ABTestSt{
		Hostport:    "http://bizsrv-cfg.srv.test.ixiaochuan.cn",
		Interval:    1,
		KeepHistory: false,
		DisableSD:   true,
	}
	err := xcabtest.Open(cfg)
	if err != nil {
		fmt.Println("err open")
		return
	}
	xcabtest.GetConfig("", "zuiyou", "", 3)
	defer xcabtest.Close()

	out := GetGiftPackParam("")
	byt, _ := json.Marshal(out)
	fmt.Println(string(byt))

	return
}
