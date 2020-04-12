package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func asyncapi() {
	fun := "asyncapi"
	fmt.Println(fun)
	urls := []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
	}
	wg := sync.WaitGroup{}
	results := []*http.Response{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			for i := 0; i < 2; i++ {
				res, err := http.Get(url)
				if err != nil {
					fmt.Println("err! url", url)
					continue
				} else {
					results = append(results, res)
					break
				}
			}
			defer wg.Done()
		}(url)
	}
	wg.Wait()
	for _, r := range results {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("fail read body. ", r.Status)
		} else {
			fmt.Println(string(content))
		}
	}
}

func syncapi() {
	fun := "syncapi"
	fmt.Println(fun)
	urls := []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
		//"http://127.0.0.1:5000",
	}
	results := []*http.Response{}
	for _, url := range urls {
		for i := 0; i < 2; i++ {
			r, err := http.Get(url)
			if err != nil {
				fmt.Println("err get. ", url)
				continue
			} else {
				results = append(results, r)
				break
			}
		}
	}
	for _, r := range results {
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("fail read body", r.Status)
		} else {
			fmt.Println(string(res))
		}
	}
}

func main() {
	fun := "main"
	st := time.Now().Unix()
	asyncapi()
	et1 := time.Now().Unix()
	syncapi()
	et2 := time.Now().Unix()
	fmt.Println(fun, et1-st, et2-et1)
}
