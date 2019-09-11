package main

import (
	//"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
)

type RecordTags struct {
	Id         string `json:"id,omitempty" bson:"_id,omitempty"`
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
	// multiple insert capable
	//err = conn.DB("testv2").C("test").Insert(rt, rt, rt)
	if err != nil {
		fmt.Println("%v", err)
	}

	// although seems ptr, it can be corretly encoded from/decoded to value
	//buf, _ := json.Marshal(rt)
	//var z xcproto.RecordTags
	//json.Unmarshal(buf, &z)
	//fmt.Println(z.GetTagid(), z.GetName())

	// getback the inserted objectid
	var result RecordTags
	err = conn.DB("testv2").C("test").Find(bson.M{
		"tagid111": 1,
	}).One(&result)
	if err == mgo.ErrNotFound {
		fmt.Println("not found tagid 1")
	} else if err != nil {
		fmt.Println("conn fail", err)
		return
	}
	if !bson.IsObjectIdHex(result.Id) {
		fmt.Println("not objectid", result.Id)
		return
	}
	objid := bson.ObjectId(result.Id).Hex()
	fmt.Println("hex: ", objid)

	var result2 RecordTags
	err = conn.DB("testv2").C("test").FindId(bson.ObjectIdHex(objid)).One(&result2)
	if err == mgo.ErrNotFound {
		fmt.Println("not found objid", objid)
	} else if err != nil {
		fmt.Println("conn fail", err)
		return
	}
	fmt.Println("r2: ", result2)

	return
}
