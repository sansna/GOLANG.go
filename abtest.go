package main

import (
	"fmt"
	"git.ixiaochuan.cn/xclib/common/xcabtest"
	"git.ixiaochuan.cn/xclib/common/xcabtest/cfgstruct"
	"strconv"
)

func main() {
	cfg := cfgstruct.ABTestSt{
		Hostport:    "http://bizsrv-cfg.srv.ixiaochuan.cn",
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

	for i := int64(0); i < 1000; i++ {
		val := strconv.FormatInt(i, 10)
		test := xcabtest.GetInt64(val, "hanabi_test_ab_v1", "test", 10)
		fmt.Println(i, test)
	}
	return
}
