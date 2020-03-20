package main

import "time"
import "fmt"

func main() {
	// get today's 0:00 ts
	ts := (time.Now().Unix()+8*3600)/86400*86400 - 8*3600
	fmt.Println(ts, time.Unix(ts, 0))

	//
	fmt.Println(time.Now().Hour())
	a := time.NewTimer(time.Second * 3)
	fmt.Println(time.Now().Unix())
	b := <-a.C
	fmt.Println(b.Unix())
	fmt.Println(time.Now().Unix())
	nanotime := time.Now().UnixNano()
	fmt.Println(nanotime)
	fmt.Printf("%020d\n", nanotime)

	var d string
	var i int64
	str := fmt.Sprintf("%025d%s", nanotime, "helloworld")
	fmt.Sscanf(str, "%025d%s", &i, &d)
	fmt.Println(i, d)
	fmt.Printf("%04d%02d%02d", time.Now().Year(), int(time.Now().Month()), time.Now().Day())

	// time ticks after waiting for 1 sec
	t := time.Tick(time.Second * 1)
	now := time.Now()
	// timeticks
	for i := range t {
		fmt.Printf("time moves %v\n", i.Second()-now.Second())
		// ensures 1 sec calls 1 time
		go func() {
			time.Sleep(time.Second * 3)
		}()
	}
}
