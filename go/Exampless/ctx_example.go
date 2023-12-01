package main

import (
	"context"
	"fmt"
	"time"
)

/*

context 是 Go 语言中用于在程序中传递截止日期、取消信号以及其他请求范围值的标准方式。它是在 Go 1.7 版本引入的，为了解决在多个 goroutine 之间传递请求范围数据的需求。

在 Go 中，每个 context 对象都可以包含一个截止日期、取消信号、请求范围值等信息。它们可以通过 context.WithDeadline、context.WithCancel、context.WithTimeout 和 context.WithValue 等函数来创建。这些函数返回一个新的 context 对象，派生自传入的 context 对象。

主要的 context 方法有：

context.Background()： 返回一个空的 context，它通常用作整个请求的根 context。

context.TODO()： 类似于 Background()，用作未来可能添加功能的预留 context。

context.WithCancel(parent)： 返回一个带有新的 Done 通道的 context，当调用返回的 cancel 函数时，会关闭该通道。

context.WithDeadline(parent, time)： 返回一个带有新的 Done 通道和截止时间的 context。当截止时间过去或调用返回的 cancel 函数时，Done 通道将关闭。

context.WithTimeout(parent, timeout)： 类似于 WithDeadline，但是截止时间是相对于当前时间的相对时间。

context.WithValue(parent, key, value)： 返回一个带有新的值的 context，该值与给定的键关联。这可以用于在请求范围内传递元数据。

使用 context 可以使得多个 goroutine 之间更容易地传递截止日期、取消信号和其他请求范围值，而无需显式传递它们作为参数。这对于处理超时、取消和传递请求范围数据等场景非常有用。
*/

func main() {
	// 创建一个根 context
	ctx := context.Background()

	// 创建一个带有取消功能的子 context 和取消函数
	ctx, cancel := context.WithCancel(ctx)

	// 启动一个 goroutine 来执行任务
	go func() {
		// 在任务完成或取消时调用 cancel 函数
		defer cancel()

		// 模拟耗时任务
		time.Sleep(3 * time.Second)

		// 在任务完成后，输出结果
		fmt.Println("Task completed")
	}()

	// 等待 context 取消
	<-ctx.Done()

	// 打印取消原因
	fmt.Println("Context canceled:", ctx.Err())
}
