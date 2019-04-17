package main

import (
	"encoding/json"
	"fmt"
	"git.ixiaochuan.cn/xclib/common/xcproto"
	goproto "github.com/gogo/protobuf/proto"

	"gopkg.in/mgo.v2/bson"

	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
	//"git.ixiaochuan.cn/gateway/service/post/record/server/logic"
)

func main() {
	testmongo, err := mongodb.InitMultiplexV2(cfgstruct.MongodbSt{
		Hostport: "127.0.0.1:27017",
		Poolsize: 16,
	}, "")
	if err != nil {
		fmt.Println("%v", err)
		return
	}

	// although seems ptr, it can be correctly saved in mongo by value.
	rt := xcproto.RecordTags{
		Tagid:      goproto.Int64(1),
		Name:       goproto.String("asdf"),
		Count:      goproto.Int64(1),
		Ct:         goproto.Int64(123),
		Cmid:       goproto.Int64(13),
		Searchable: goproto.Int64(0),
		Status:     goproto.Int64(0),
	}

	conn := testmongo.GetConn()
	//for i := 0; i < 100; i++ {
	//	err = conn.DB("testv2").C("test").Insert(rt)
	//	if err != nil {
	//		fmt.Println("%v", err)
	//	}
	//}

	query := bson.M{}
	result := []*xcproto.RecordTags{}
	// mgo.v2's find returns all elem.
	err = conn.DB("testv2").C("test").Find(query).All(&result)
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Println(len(result), "elem in testv2")

	// although seems ptr, it can be corretly encoded from/decoded to value
	buf, _ := json.Marshal(rt)
	var z xcproto.RecordTags
	json.Unmarshal(buf, &z)
	fmt.Println(z.GetTagid(), z.GetName())

	return
}
