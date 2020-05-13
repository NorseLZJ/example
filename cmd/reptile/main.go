package main

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/NorseLZJ/example/std"
	"github.com/NorseLZJ/example/std/cfg_marshal"
	"github.com/NorseLZJ/example/std/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	share = flag.String("f", "./reptile.json", "cfg")
	cfgT  = &cfg_marshal.Reptile{}
	db    *gorm.DB
	llog  *log.Log

	soldUrl   string
	sellUrl   string
	userAgent string
)

const (
	spaceTime = 1
)

func init() {
	llog = log.LLog()
}

func checkCfg() {
	if cfgT.Url.SellUrl == "" || cfgT.Url.SoldUrl == "" {
		std.CheckErr(errors.New("sell url or sold url is nil"))
	}
	sellUrl = cfgT.Url.SellUrl
	soldUrl = cfgT.Url.SoldUrl
	userAgent = cfgT.UserAgent
}

func main() {
	flag.Parse()
	err := cfg_marshal.Marshal(*share, cfgT)
	std.CheckErr(err)
	checkCfg()
	initDb()

	defer func() {
		db.Close()
	}()

	district := []string{"weiyang", "yanta"}
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for _, val := range district {
			total := sellingPage(val)
			for page := 1; page < total; page++ {
				wg.Add(1)
				time.Sleep(time.Duration(spaceTime) * time.Second)
				go func(page int) {
					defer wg.Done()
					sellingInfo(val, page)
				}(page)
			}
		}
	}()

	go func() {
		for _, val := range district {
			total := soldPage(val)
			for page := 1; page < total; page++ {
				wg.Add(1)
				time.Sleep(time.Duration(spaceTime) * time.Second)
				go func(page int) {
					defer wg.Done()
					soldInfo(val, page)
				}(page)
			}
		}
	}()
	wg.Wait()
}

func initDb() {
	var err error
	driverName := cfgT.SqlConfig.Driver
	host := cfgT.SqlConfig.Host
	port := cfgT.SqlConfig.Port
	database := cfgT.SqlConfig.Db
	username := cfgT.SqlConfig.User
	password := cfgT.SqlConfig.Password
	charset := cfgT.SqlConfig.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err = gorm.Open(driverName, args)
	std.CheckErr(err)
	db = db.AutoMigrate(&Selling{}, &Sold{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
}
