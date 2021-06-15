package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	n := make(chan string, 10000)
	end := make(chan bool)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	defer wg.Wait()
	go dispose(1, n, end, wg)
	go dispose(2, n, end, wg)
	go dispose(3, n, end, wg)
	go dispose(4, n, end, wg)
	for i := 0; i < 100; i++ {
		n <- "hello"
	}
	for {
		time.Sleep(time.Second)
		if len(n) <= 0 {
			end <- true
			end <- true
			end <- true
			end <- true
			return
		}
	}
	time.Sleep(1)
}

func WalkDir(dirPth, suffix string, out chan string) {
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err := filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			out <- filename
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func dispose(id int, in chan string, end chan bool, done *sync.WaitGroup) {
	defer done.Done()
	for {
		select {
		case k, ok := <-in:
			if ok {
				fmt.Printf("%d\n", k)
			}
		case b := <-end:
			_ = b
			return
		}
	}
}
