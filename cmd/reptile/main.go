package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Reptile struct {
	// sql
	SqlConfig struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
	} `json:"sqlConfig"`
	// url
	Url struct {
		SoldUrl string `json:"sold_url"`
		SellUrl string `json:"sell_url"`
	} `json:"url"`
	// userAgent
	UserAgent string `json:"user_agent"`
	// 区
	District []string `json:"district"`
}

var (
	share = flag.String("f", "./reptile.json", "cfg")
	cfgT  = &Reptile{}
	db    *gorm.DB

	soldUrl    string
	sellUrl    string
	userAgent  string
	city       string
	district   []string
	createTime = time.Now()
)

const (
	spaceTime = 1
)

func checkCfg() {
	if cfgT.Url.SellUrl == "" || cfgT.Url.SoldUrl == "" {
		panic(errors.New("sell url or sold url is nil"))
	}
	sellUrl = cfgT.Url.SellUrl
	soldUrl = cfgT.Url.SoldUrl
	userAgent = cfgT.UserAgent
	strs := strings.Split(sellUrl, "/")
	city = strings.Split(strs[2], ".")[0]
	district = make([]string, 0, len(cfgT.District))
	district = append(district, cfgT.District[:]...)
}

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*share)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, cfgT)
	if err != nil {
		panic(err)
	}
	checkCfg()
	initDb()

	defer func() {
		db.Close()
	}()
	wg := &sync.WaitGroup{}
	sell(wg)
	sold(wg)
	wg.Wait()
}

func sell(wg *sync.WaitGroup) {
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
}

func sold(wg *sync.WaitGroup) {
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
	if err != nil {
		panic(err)
	}
	db = db.AutoMigrate(&Selling{}, &Sold{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
}
