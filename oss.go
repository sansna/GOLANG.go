package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-audio/wav"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"os/exec"
	"strings"

	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
)

var (
	Acntmgo   *mongodb.MultiplexClient
	Attmgo    *mongodb.MultiplexClient
	Yonmgo    *mongodb.MultiplexClient
	Postmgo   *mongodb.MultiplexClient
	Recordmgo *mongodb.MultiplexClient
	Matchmgo  *mongodb.MultiplexClient
	Revmgo    *mongodb.MultiplexClient
	Chatmgo   *mongodb.MultiplexClient
	Mediamgo  *mongodb.MultiplexClient
	Oprdwmgo  *mongodb.MultiplexClient
	Oprmgo    *mongodb.MultiplexClient
)

type Result struct {
	Objid    int64 `bson:"objid"`
	Notified int64 `bson:"notified"`
}

func main() {
	//init
	ossapi, err := oss.New("http://oss-cn-shanghai-internal.aliyuncs.com", "", "")
	if err != nil {
		fmt.Println("fail new.")
		return
	}
	bucket, err := ossapi.Bucket("hanabi")
	if err != nil {
		fmt.Println("fail new.")
		return
	}

	Oprmgo, err := mongodb.InitMultiplexV2(cfgstruct.MongodbSt{
		Hostport: "",
		UserName: "",
		Password: "",
		Poolsize: 8,
	}, "")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	marker := oss.Marker("")
	for {
		lsres, err := bucket.ListObjects(oss.MaxKeys(200), marker)
		if err != nil {
			fmt.Println("fail ls.")
			return
		}
		for _, item := range lsres.Objects {
			filename := item.Key
			sid_cname := strings.Split(strings.Split(filename, ".")[0], "_")
			if len(sid_cname) != 2 {
				fmt.Println("err filename")
				continue
			}
			sid := sid_cname[0]
			cname := sid_cname[1]
			if len(sid) == 0 || len(cname) == 0 {
				fmt.Println("fail get cname.")
				continue
		}
			err = bucket.GetObjectToFile(filename, "./musics/"+filename)
			if err != nil {
				fmt.Println("fail down ", filename)
				continue
			}
			// get audio dur
			f, _ := os.Open("./musics/" + filename)
			w := wav.NewDecoder(f)
			dur, err := w.Duration()
			f.Close()
			cmd := exec.Command("/usr/bin/bash", "-c", "rm -frd ", "./musics/"+filename)
			cmd.Run()
			if err != nil {
				fmt.Println(err, "fail get dur.")
				continue
			}
			// get reported, notified from mongo.
			result := &Result{}
			err = Oprmgo.FindOne("", "", "", bson.M{"channel_name": cname}, 0, result)
			if err == mgo.ErrNotFound {
				err = nil
				fmt.Println("not found")
			} else if err != nil {
				fmt.Println(" fail found.")
				continue
			}
			fmt.Println(dur.Seconds(), result.Objid, result.Notified)
		}
		marker = oss.Marker(lsres.NextMarker)
		// Fin
		if !lsres.IsTruncated {
			break
		}
	}
}
