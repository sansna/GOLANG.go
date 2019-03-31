package main

import (
	//"encoding/json"
	"fmt"

	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
)

type RecordTags struct {
	Tagid      int64  `json:"tagid" bson:"tagid111"`
	Name       string `json:"name"`
	Count      int64  `json:"asdfcount" bson:"asdfxzn"`
	Ct         int64  `json:"ct"`
	Cmid       int64  `json:"cmid"`
	Searchable int64  `json:"searchable"`
	Status     int64  `json:"status"`
}

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
	rt := RecordTags{
		Tagid:      1,
		Name:       "asdf",
		Count:      1,
		Ct:         123,
		Cmid:       13,
		Searchable: 0,
		Status:     0,
	}

	conn := testmongo.GetConn()
	err = conn.DB("testv2").C("test").Insert(rt)
	if err != nil {
		fmt.Println("%v", err)
	}

	// although seems ptr, it can be corretly encoded from/decoded to value
	//buf, _ := json.Marshal(rt)
	//var z xcproto.RecordTags
	//json.Unmarshal(buf, &z)
	//fmt.Println(z.GetTagid(), z.GetName())

	return
}
