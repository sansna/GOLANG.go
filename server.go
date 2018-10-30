package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"io"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

var (
	Redisserver = "localhost:6379"
	Mgoserver = "localhost"
)

var (
	NanoSec = time.Duration(1)
	MiuSec = 1000*NanoSec
	MiliSec = MiuSec*1000
	Second = MiliSec*1000
	Minute = Second*60
	Hour = Minute*60
	Day = Hour*24
)

type Keyvalue struct {
	Key string `bson:"k"`
	Val string `bson:"v"`
}

type Keys struct {
	Key []string `bson:"ks"`
}

func sendresp(w http.ResponseWriter, errno int, msg string) {
	if (errno == -1) {
		fmt.Fprintf(w,"{\"ret\":%d, \"msg\":\"%s\"}", errno, msg)
	} else if (errno == 1) {
		fmt.Fprintf(w,"{\"ret\":%d, \"data\":\"%s\"}", errno, msg)
	}
}

func redis_conn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: Redisserver,
		Password: "",
		DB: 0,
	})

	err := client.Ping().Err()
	if (err != nil) {
		panic(err)
	}
	//defer client.Close()
	
	return client
}

func mgo_conn() *mgo.Session {
	session, err := mgo.Dial(Mgoserver)
	if (err != nil) {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	return session
}

func addhandler(w http.ResponseWriter, r *http.Request) {
	session := mgo_conn()
	defer session.Close()

	switch(r.Method) {
	case "POST":
		var t Keyvalue
		var data map[string]interface{}
		data = make(map[string]interface{})

		dcd := json.NewDecoder(r.Body)
		for {
			err := dcd.Decode(&t)
			if (err == io.EOF) {
				break
			} else if (err != nil) {
				panic(err)
			}
			err = session.DB("main").C("kv").Insert(&Keyvalue{t.Key,t.Val})
			if (err != nil) {
				panic(err)
			}
			data[t.Key] = t.Val
		}
		ret, err := json.Marshal(data)
		if (err != nil) {
			panic(err)
		}
		sendresp(w, 1, string(ret))
	default:
		sendresp(w, -1, "only POST supported.")
	}
}

func delhandler(w http.ResponseWriter, r *http.Request) {
	session := mgo_conn()
	defer session.Close()

	switch(r.Method) {
	case "POST":
		//format problem
		dcd := json.NewDecoder(r.Body)
		var ret strings.Builder
		var t Keys
		err := dcd.Decode(&t)
		if (err != io.EOF && err != nil) {
			panic(err)
		}
		client := redis_conn()
		defer client.Close()

		for i := 0; i < len(t.Key); i++ {
			// remove from mgo
			err = session.DB("main").C("kv").Remove(bson.M{"k":t.Key[i]})
			if (err != nil) {
				//XXX: not found, ensure partial case
				//sendresp(w, -1, "Not found in db.")
				//return
				//panic(err)
				continue
			}
			// del from redis
			err = client.Del(t.Key[i]).Err()
			if (err != nil) {
				panic(err)
			}
			ret.WriteString("{\"Key\":\"")
			ret.WriteString(t.Key[i])
			ret.WriteString("\"}")
		}
		if (ret.Len() > 0) {
			sendresp(w, 1, ret.String())
		} else {
			sendresp(w, -1, "Not found any in db.")
		}

	default:
		sendresp(w, -1, "only POST supported.")
	}
}

func modhandler(w http.ResponseWriter, r *http.Request) {
	session := mgo_conn()
	defer session.Close()

	switch(r.Method) {
	case "POST":
		var ret strings.Builder
		var t Keyvalue
		var count int
		count = 0
		dcd := json.NewDecoder(r.Body)
		client := redis_conn()
		defer client.Close()

		for {
			err := dcd.Decode(&t)
			if (err == io.EOF) {
				break
			} else if (err != nil) {
				panic(err)
			}
			err = session.DB("main").C("kv").Update(bson.M{"k":t.Key}, bson.M{"$set": bson.M{"v":t.Val}})
			if (err != nil) {
				// ensure partial case
				continue
				//sendresp(w, -1, "Cannot mod value, because does not exist.")
				//return
			}
			count++
			err = client.Del(t.Key).Err()
			if (err != nil) {
				// redis timeout, ignore case
				//panic(err)
			}
			ret.WriteString("{\"Key\":\"")
			ret.WriteString(t.Key)
			ret.WriteString("\",\"Val\":\"")
			ret.WriteString(t.Val)
			ret.WriteString("\"}")
		}
		if (count > 0) {
			sendresp(w, 1, ret.String())
		} else {
			sendresp(w, -1, "Not found any in db.")
		}

	default:
		sendresp(w, -1, "only POST supported.")
	}
}

func quehandler(w http.ResponseWriter, r *http.Request) {
	switch (r.Method) {
	case "POST":
		var t Keys
		var ret strings.Builder
		dcd := json.NewDecoder(r.Body)
		err := dcd.Decode(&t)
		if (err != nil) {
			panic(err)
		}
		if (len(t.Key) <= 0) {
			return
		}
		session := mgo_conn()
		defer session.Close()
		client := redis_conn()
		defer client.Close()

		var data map[string]interface{}
		data = make(map[string]interface{})
		for i:=0; i < len(t.Key); i++ {
			val, err := client.Get(t.Key[i]).Result()
			//XXX: determine get failure errnos
			if (err != nil) {
				// try from mgo
				var value Keyvalue
				err = session.DB("main").C("kv").Find(bson.M{"k":t.Key[i]}).One(&value)
				if (err != nil) {
					//XXX: not found in mgo; ensure partial case
					continue
					//sendresp(w, -1, "Not found in db.")
					//return
				}
				val = value.Val
				//push back to redis
				err = client.Set(value.Key, value.Val, 20*Second).Err()
				if (err != nil) {
					panic(err)
				}
			}
			data[t.Key[i]] = val
		}
		newdata, err := json.Marshal(data)
		if (err != nil) {
			panic(err)
		}
		ret.WriteString(string(newdata))
		sendresp(w, 1, ret.String())
	default:
		sendresp(w, -1, "only POST supported.")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", addhandler)
	// consider partial case
	mux.HandleFunc("/mod", modhandler)
	// consider partial case
	mux.HandleFunc("/del", delhandler)
	// consider partial case
	mux.HandleFunc("/query", quehandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
