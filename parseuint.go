package main

import (
	"encoding/json"
	"fmt"

	//"git.ixiaochuan.cn/gateway/ctrl"
	"git.ixiaochuan.cn/xclib/adapter/proto/message/proto"
)

func main() {
	p := &proto.MDelSessionsParam{}
	//c := &ctrl.Controller{}
	// Note: msgid/session_id both uint64 in protobuf. Thus cannot be decoded as negative int.
	err := json.Unmarshal([]byte("{\"sessions\":[{\"msgid\":\"-81723\",\"session_id\":\"-6805855164747448323\"}]}"), p)
	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println(p.GetMid(), p.GetToken(), p.GetSessions()[0].GetSessionId())
	}
}
