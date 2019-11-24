package main

import (
	"fmt"
	"git.ixiaochuan.cn/xclib/common/xcabtest"
	"git.ixiaochuan.cn/xclib/common/xcabtest/cfgstruct"
)

func main() {
	xcabtest.Open(cfgstruct.ABTestSt{
		Hostport:    "http://bizsrv-cfg.srv.ixiaochuan.cn",
		Interval:    1,
		KeepHistory: false,
		DisableSD:   true,
	})
	defer xcabtest.Close()
	var ids []int64
	retids := xcabtest.GetInt64Slice("b0255008af9f2ea3", "hh_chat_accompany", "mids", ids)
	fmt.Println(ids, retids)
}
