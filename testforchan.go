package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 10)
	ch <- true
	ch <- false
	ch <- true
	go func() {
		time.Sleep(time.Second * 3)
		close(ch)
	}()
	for {
		select {
		case val, ok := <-ch:
			// ok: 检测是否channel关闭
			// channel关闭后，如果不检测channel是否关闭,select总是会执行本case,而不会走default
			if ok {
				fmt.Println("ok", val)
			} else {
				// channel关闭时的val是个默认值
				fmt.Println("fail", val)
				return
			}
		default:
			// without this default, the program is like blocked.
			time.Sleep(time.Millisecond)
		}
	}

}
