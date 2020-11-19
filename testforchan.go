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
	breakFor := false
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
				breakFor = true
				break
			}
		default:
			// without this default, the program is like blocked.
			time.Sleep(time.Millisecond)
		}
		if breakFor {
			break
		}
	}

	//
	// one cannot reopen a channel
	ch = make(chan bool, 4)
	go func() {
		ch <- false
		ch <- true
		ch <- false
		time.Sleep(time.Second)
		close(ch)
	}()
	for c := range ch {
		fmt.Println(c)
	}
}
