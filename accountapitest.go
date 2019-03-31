package main

import (
	"context"
	"fmt"
	"git.ixiaochuan.cn/xclib/adapter/api/accountapi"
	"git.ixiaochuan.cn/xclib/adapter/api/cfgstruct"
)

func main() {
	var mid int64
	mid = 62885756
	cfg := cfgstruct.AccountApiCfgSt{
		ApiCfgSingleSt: cfgstruct.ApiCfgSingleSt{
			AppHost: "http://acnt.srv.in.ixiaochuan.cn",
		},
	}
	accountapi.Init(cfg)
	resp, _ := accountapi.New(context.Background()).GetMemberByMid(mid)
	//list := resp.GetMedal()
	//fmt.Println(len(list))
	fmt.Println(*resp.IsReg)
	fmt.Println(resp.GetProfession())
}
