package main

import (
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	maxProcess = 8
)

func main() {
	num := os.Args[1]
	switch num {
	case "1":
		runtime.GOMAXPROCS(1)
	case "2":
		runtime.GOMAXPROCS(2)
	case "4":
		runtime.GOMAXPROCS(4)
	case "8":
		runtime.GOMAXPROCS(8)
	}

	n := make(chan string, 10000)
	end := make(chan bool, maxProcess)

	wg := &sync.WaitGroup{}
	wg.Add(maxProcess)
	for i := 1; i <= maxProcess; i++ {
		go dispose(i, n, end, wg)
	}

	for i := 0; i < 10000000; i++ {
		n <- "hello1"
		n <- "hello2"
		n <- "hello3"
		n <- "hello4"
		n <- "hello5"
		n <- "hello6"
		n <- "hello7"
		n <- "hello8"
		n <- "hello9"
		n <- "hello10"
	}
	for {
		time.Sleep(time.Second)
		if len(n) <= 0 {
			for i := 0; i < maxProcess; i++ {
				end <- true
			}
			return
		}
	}
	wg.Wait()
}

func dispose(id int, in chan string, end chan bool, done *sync.WaitGroup) {
	defer done.Done()
	for {
		select {
		case k, ok := <-in:
			if ok {
				//fmt.Printf("from:%d ->  %v\n", id, k)
				_ = k
			}
		case b := <-end:
			_ = b
			return
		}
	}
}
