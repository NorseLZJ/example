package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 启动一个 goroutine 来执行任务
	go func() {
		// 在任务完成或超时时，自动触发 context 取消
		// 这里模拟一个耗时的任务，超过超时时间
		fmt.Println("sleep")
		time.Sleep(3 * time.Second)
		fmt.Println("week")
	}()

	// 等待 context 取消
	<-ctx.Done()

	// 打印取消原因
	fmt.Println("Context canceled:", ctx.Err())
}
