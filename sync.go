// wg is used as pthread_join
// sync is used as pthread_mutex
package main

import (
	"fmt"
	"sync"
	//	"time"
)

type SyncData struct {
	*sync.RWMutex
	a *int64
}

type SyncDatas struct {
	sync.RWMutex
	sd map[int64]SyncData
}

func (a *SyncDatas) thread(n int64, wg *sync.WaitGroup) {
	a.Lock()
	if c, ok := a.sd[1]; !ok {
		c.RWMutex = &sync.RWMutex{}
		c.a = new(int64)
		a.sd[1] = c
	}
	a.Unlock()
	a.sd[1].Lock()
	*a.sd[1].a = *a.sd[1].a + n
	a.sd[1].Unlock()
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	// only when passing a as address works.
	a := &SyncDatas{}
	a.sd = make(map[int64]SyncData, 1)
	i := int64(0)
	for ; i < 10000; i++ {
		wg.Add(1)
		go a.thread(i, &wg)
	}
	wg.Wait()
	fmt.Println(*a.sd[1].a)
}
