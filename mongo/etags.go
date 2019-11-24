package main

import (
	"fmt"
	"strings"

	"github.com/wangbin/jiebago/analyse"
	"gopkg.in/mgo.v2/bson"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
	"git.ixiaochuan.cn/xclib/common/xcproto"
	rds_cfg "git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

func main() {
	acntmgo, err := mongodb.InitMultiplexV2(cfgstruct.MongodbSt{
		Hostport: "dds-uf65797e0d7307141.mongodb.rds.aliyuncs.com:3717",
		UserName: "hanabi",
		Password: "Dei7cai4uYookoo",
		Poolsize: 8,
	}, "account")
	if err != nil {
		fmt.Println("err: %v", err)
		return
	}
	cache, err := base.NewCacheV2(rds_cfg.RedisSt{
		Hostport: "r-uf6934ae6d85c604.redis.rds.aliyuncs.com:6379",
		Poolsize: 8,
		Timeout:  30,
		Auth:     "crypt-124e9177ac8b466ee3c8c6713cf7a8f06704ed13952238941ba9ecd1db83f840",
	}, "account-redis")

	var te analyse.TagExtracter
	err = te.LoadDictionary("/app/golang/online_hanabi-bizsrv-acnt/etc/dict.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = te.LoadIdf("/app/golang/online_hanabi-bizsrv-acnt/etc/dict.txt")
	if err != nil {
		fmt.Println(err)
	}

	profst := make([]*xcproto.ProfileSt, 0)
	query := bson.M{
		"tag_list": bson.M{
			"$exists": 1,
		},
		"e_tagid_list": bson.M{
			"$exists": 1,
		},
	}
	err = acntmgo.Find("account", "member_profile", "", query, 0, 0, &profst)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	var str string
	for _, p := range profst {
		taglist := p.GetTagList()
		//fmt.Println("prf", p.GetMid(), taglist)
		strtags := make([]string, 0)
		for _, tag := range taglist {
			strtags = append(strtags, tag.GetTagName())
		}
		str = strings.Join(strtags, " ")
		etags := te.ExtractTags(str, 20)
		strs := make([]string, 0)
		for _, etag := range etags {
			strs = append(strs, etag.Text())
		}

		tags := make([]*xcproto.ETagItem, 0)
		err = acntmgo.Find("account", "e_tags", "", bson.M{
			"e_tag_name": bson.M{
				"$in": strs,
			},
		}, 0, 0, &tags)
		if err != nil {
			fmt.Println("err etags err:", err)
			return
		}

		// reorder
		ids := make([]int64, 0)
		for _, s := range strs {
			for _, t := range tags {
				if t.GetETagName() == s {
					ids = append(ids, t.GetETid())
					break
				}
			}
		}
		err = acntmgo.Update("account", "member_profile", bson.M{
			"_id": p.GetMid(),
		}, bson.M{
			"$set": bson.M{
				"e_tagid_list": ids,
			},
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		key := fmt.Sprintf("bizacnt_member-profile-%d", p.GetMid())
		cache.Del(key)
	}
}
