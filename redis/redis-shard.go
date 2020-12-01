package main

import "math/rand"

//import "sort"
import "sync"
import "strconv"
import "fmt"
import "time"
import "github.com/go-redis/redis"

const (
	m      = 2500
	n      = 10000
	minmid = 1
	maxmid = 1000000
)

type A struct {
	V *[n]int64
}

func (a A) Less(i, j int) bool {
	return a.V[i] > a.V[j]
}

func (a A) Len() int {
	return n
}

func (a A) Swap(i, j int) {
	a.V[i], a.V[j] = a.V[j], a.V[i]
}

func InSlice(s []int64, in int64) bool {
	for _, v := range s {
		if in == v {
			return true
		}
	}
	return false
}

type Info struct {
	Idx  int64
	Mids []string
}

func GetRandMids(count int64) []int64 {
	rand.Seed(time.Now().UnixNano())
	out := []int64{}
	for i := int64(0); i < count; i++ {
		genmid := int64(rand.Intn(int(maxmid-minmid)) + int(minmid))
		if InSlice(out, genmid) {
			i--
			continue
		}
		out = append(out, genmid)
	}
	return out
}

func TransSliceIntToSliceInterface(in []int64) (out []interface{}) {
	out = []interface{}{}
	for _, v := range in {
		out = append(out, v)
	}
	return
}

func InitRedis(port string) *redis.Client {
	if len(port) == 0 {
		port = "6379"
	}
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: "",
		DB:       0,
	})
	return r
}

func ResetVal(rds *redis.Client, key string, mems []interface{}) {
	rds.Del(key)
	rds.SAdd(key, mems...)
}

func InitData(rds, rds2 *redis.Client) {
	friends := GetRandMids(2000)
	liked := GetRandMids(10000)
	ResetVal(rds, "mfset", TransSliceIntToSliceInterface(friends))
	ResetVal(rds2, "mfset", TransSliceIntToSliceInterface(friends))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset1", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset2", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset3", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset4", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset5", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset6", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset7", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset8", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset9", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset10", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset11", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset12", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset13", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset14", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset15", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset16", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset17", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset18", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds2, "plset19", TransSliceIntToSliceInterface(liked))
	liked = GetRandMids(10000)
	ResetVal(rds, "plset20", TransSliceIntToSliceInterface(liked))
}

func main() {
	rds := InitRedis("6379")
	rds2 := InitRedis("6370")
	//rds.Set("a", 100, 0)
	//val := rds.Get("a")
	//fmt.Println(val)

	InitData(rds, rds2)

	var plnames []string
	for i := int64(1); i < 21; i++ {
		plname := "plset" + strconv.FormatInt(i, 10)
		plnames = append(plnames, plname)
	}
	st := time.Now().UnixNano()
	ets := []int64{}
	ch := make(chan Info, 30)

	wg := sync.WaitGroup{}
	for i := int64(1); i < 21; i++ {
		wg.Add(1)
		idx := i
		go func() {
			// NOTE: redis is single threaded. the more the redis instance, the faster
			// The more usage with redis, the more potential to fail. So try
			// with context timeout and with default value.
			if idx%2 == 1 {
				res, _ := rds2.SInter("mfset", plnames[idx-1]).Result()
				ch <- Info{
					Idx:  idx,
					Mids: res,
				}
			} else {
				res, _ := rds.SInter("mfset", plnames[idx-1]).Result()
				ch <- Info{
					Idx:  idx,
					Mids: res,
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)

	ets = append(ets, time.Now().UnixNano()-st)
	//fmt.Println(list.Result())
	fmt.Println(ets)

	for info := range ch {
		fmt.Println(info)
	}

	return
	//st := time.Now().UnixNano()
	//ets := []int64{}
	//rand.Seed(time.Now().UnixNano())
	//x := [n]int64{}

	//for i := 0; i < m; i++ {
	//	r := int64(rand.Intn(n))
	//	x[r]++
	//}
	//ets = append(ets, time.Now().UnixNano()-st)

	//a := A{
	//	V: &x,
	//}

	//sort.Sort(a)
	////fmt.Println(*a.V)
	//ets = append(ets, time.Now().UnixNano()-st)

	//hit_count := 0
	//for _, v := range a.V {
	//	if v > 0 {
	//		hit_count++
	//	}
	//}
	//ets = append(ets, time.Now().UnixNano()-st)

	//fmt.Println("hitcount: ", hit_count, "ets:", ets)
}
