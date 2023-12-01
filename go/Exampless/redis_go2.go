package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func redisIsOk(a interface{}, b error) bool {
	if a == nil {
		return false
	}
	switch a.(type) {
	case string:
		x := a.(string)
		return "OK" == x
	case int64:
		x := a.(int64)
		return 1 == x
	}
	return false
}

func setKey(index int, wg *sync.WaitGroup) {
	conn := pool.Get()

	defer func() {
		_ = conn.Close()
		wg.Done()
	}()

	var (
		key = "lock_shop_1"
		//value = 1
	)

	b := redisIsOk(conn.Do("set", key, index, "ex", 60, "nx"))
	if !b {
		log.Printf("用户 %d 号 设置失败 \n", index)
		return
	}

	fmt.Printf("用户 %d 号 正在购买 \n", index)
	time.Sleep(time.Second * 2)
	//if b = redisIsOk(conn.Do("DEL", key)); b {
	//	fmt.Printf("用户 %d 号 购买结束 \n", index)
	//	return
	//} else {
	//	fmt.Printf("用户 %d 号 购买结束 !!!! \n", index)
	//}
}

func main() {

	// 创建 Redis 连接池
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	var wg sync.WaitGroup
	numUsers := 5

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go setKey(i, &wg)
	}
	time.Sleep(time.Second * 5)

	// 等待所有用户完成设置
	wg.Wait()
}
