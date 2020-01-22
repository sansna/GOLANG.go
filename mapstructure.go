package main

import (
	"fmt"
	"git.ixiaochuan.cn/xclib/adapter/proto/spam/proto"
	"github.com/mitchellh/mapstructure"
)

func main() {
	s := &proto.HanabiBizspamOptionsWrap{}
	a := make(map[string]interface{})
	a["conf"] = map[string]interface{}{
		"blockconf": 1, // BlockConf structure mapped, failed, though.
	}
	err := mapstructure.Decode(a, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s, s.Conf)
	return
}
