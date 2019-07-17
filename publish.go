package main

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"git.ixiaochuan.cn/xclib/adapter/proto/stickers/proto"
	"git.ixiaochuan.cn/xclib/common/xcproto"
	goproto "github.com/gogo/protobuf/proto"
)

func Publish(url, msg string) (resp *http.Response, err error) {
	resp, err = http.Post(url, msg, nil)
	return resp, err
}

func main() {
	msg := &proto.GetStickersParam{
		Mid: goproto.Int64(1),
		Sid: goproto.Int64(1),
	}
	codec := &xcproto.ProtoCodec{}
	byt, err := codec.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, _ := Publish("172.20.20.89:9009/stickers/httpapi/get_stickers", string(byt))
	fmt.Println(resp)
}
