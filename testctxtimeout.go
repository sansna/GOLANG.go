package main

import (
	"context"
	"fmt"
	"time"

	"git.ixiaochuan.cn/pp_server/live/proto/adapter/miscapi"
	"git.ixiaochuan.cn/xc_server/xclib/common/cconf"
	"git.ixiaochuan.cn/xc_server/xclib/common/jobqueue"
)

func main() {
	jobqueue.SetUsingNewClient(true)
	cconf.InitIniConfer("localhost:8500", "ppapiconfs")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := make(chan interface{}, 10)

	go func() {
		resp, err := miscapi.New(ctx).TestTimeout(miscapi.TestTimeoutParam{
			Ms: 500,
		})
		if err != nil {
			fmt.Println("500 no", err)
		} else {
			fmt.Println(resp)
			ch <- resp
		}
	}()

	select {
	case <-ctx.Done():
		// 到期
		fmt.Println("daoqi")
		return
	case val := <-ch:
		// 请求按时返回结果
		vals := val.(miscapi.TestTimeoutResp)
		fmt.Println(vals)
		return
	}
}
