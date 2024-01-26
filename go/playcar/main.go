package main

import (
	"flag"
	"log"
	"playcar/dbm"
	rdb "playcar/rd_m"
	"playcar/route"
	"runtime"
	"time"
)

var (
	config = flag.String("config", "config.json", "server config file")
)

func main() {
	flag.Parse()
	ConfInit()
	dbm.Init(GConfig.Mysql)
	rdb.InitLeaderboard(GConfig.Redis.Host, GConfig.Redis.Password, GConfig.Redis.Db)
	dbm.LoadAllUser()
	go printNumGoroutinePeriodically(time.Millisecond * 500)
	// gin 会自己阻塞
	route.Init(GConfig.JwtKey, GConfig.Listen)
}

func printNumGoroutinePeriodically(interval time.Duration) {
	num, prevNum := 0, 0
	for {
		num = runtime.NumGoroutine()
		if num != prevNum {
			log.Printf("Go routine num:%d\n", num)
			prevNum = num
		} else {
			time.Sleep(interval)
		}
	}
}
