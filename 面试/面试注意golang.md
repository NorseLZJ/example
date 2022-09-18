
# 一些笔试题链接

https://jishuin.proginn.com/p/763bfbd4cfd5


# 需要仔细看看的问题

# select

-  一般用法

```go
package main

import "fmt"

func send(out chan int) {
	fmt.Println("send1")
	out <- 1
	fmt.Println("send2")
	close(out)
}

func main() {
	out := make(chan int)
	go send(out)
	for {
		select {
		case v, ok := <-out:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("chan closed")
				return
			}
			break
		default:
			continue
		}
	}
}

/*
run result:
send1
1
send2
chan closed
*/

```

# channel

- channel 就是一个加锁的,环形(队列|数组) version:1.19.1 chan 定义 
```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements 
	// 存数据的数组
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index    
	// 发送到哪一个下标
	recvx    uint   // receive index 
	// 存储到哪一个下标
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
	// 锁,保证多线程安全
}

```

# 注意点

- 如果一个结构里边有map,优先考虑,map没有初始化就使用

- map 的使用必须初始化, 大括号初始化,或者make

- 函数定义和传参数必须要一致,接收指针,或者接收变量要看清

- 如果函数参数有interface{} ,参数有断言,必须判断是否成功,不然会有可能panic


# 特殊代码

```go
func main() {
	// slice ,map 如果没有用两个变量接收,都取的是index,
	x := []string{"a", "b", "c"}
	for _, v := range x {
		fmt.Print(v)
	}
}

/*---------------------*/
func main() {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}
// result:
// three three three

/*---------------------*/
type Slice []int
func NewSlice() Slice {
         return make(Slice, 0)
}
func (s* Slice) Add(elem int) *Slice {
         *s = append(*s, elem)
         fmt.Print(elem)
         return s
}
func main() {  
         s := NewSlice()
         defer s.Add(1).Add(2).Add(3)
         s.Add(4)
}
// result:1243 
// 最后一条语句总是在倒数第二个位置出现,我也不知道为什么

```