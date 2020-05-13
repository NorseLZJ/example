package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NorseLZJ/example/std"
	"github.com/gocolly/colly"
)

type Selling struct {
	Id         string `gorm:"varchar(64) ;primary_key ;comment: '房子id'"`
	Name       string `gorm:"varchar(64);comment:'小区名称'"`
	TotalPrice int    `gorm:"comment: '房子总价'"`
	UnitPrice  int    `gorm:"comment: '房子单价'"`
	District   string `gorm:"varchar(64); comment:'所属行政区'"`
	Region     string `gorm:"varchar(64); comment:'详细区域'"`
	Area       int    `gorm:"comment:'面积'"`
}

type Sold struct {
	Id         string `gorm:"varchar(64) ;primary_key ;comment: '房子id'"`
	Name       string `gorm:"varchar(64);comment:'小区名称'"`
	TotalPrice int    `gorm:"comment: '房子总价'"`
	UnitPrice  int    `gorm:"comment: '房子单价'"`
	District   string `gorm:"varchar(64); comment:'所属行政区'"`
	SoldYear   string `gorm:"varchar(32)  ;comment: '交易年份'"`
	SoldMonth  string `gorm:"varchar(32)  ;comment: '交易月份'"`
	Area       int    `gorm:"comment:'面积'"`
}

func soldInfo(community string, page int) {
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent(userAgent),
	)
	c.SetRequestTimeout(time.Duration(120) * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: soldUrl, Parallelism: 1}) //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
		llog.Info("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
	})
	//访问所有info 访问前20页采用goroutine
	c.OnHTML(".listContent>li", func(e *colly.HTMLElement) {
		re, _ := regexp.Compile(`\d+`)                                                                             //正则表达式用来匹配数字
		houseId := string(re.Find([]byte(strings.Split(e.ChildAttr("div.info > div.title > a", "href"), "/")[4]))) //获取房子ID，可根据ID直接访问房子详情主页
		name := strings.Split(e.ChildText("div.info > div.title > a"), " ")[0]                                     //获取小区名
		area := 0
		if len(strings.Split(e.ChildText("div.info > div.title > a"), " ")) == 3 {
			area, _ = strconv.Atoi(string(re.Find([]byte(strings.Split(e.ChildText("div.info > div.title > a"), " ")[2])))) //获取总面积
		}
		totalPrice, _ := strconv.Atoi(e.DOM.Find(".info .address .totalPrice span").Eq(0).Text())                      //获取总价
		unitPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .flood .unitPrice span").Eq(0).Text())))) //获取单价
		dealDate := e.DOM.Find(".info .address .dealDate").Eq(0).Text()                                                //获取成交年月日
		soldYear := strings.Split(dealDate, ".")[0]                                                                    //分离出成交年份
		soldMonth := strings.Split(dealDate, ".")[1]                                                                   //分离出成交月
		if houseId != "" {
			sold := Sold{
				Id:         houseId,
				Name:       name,
				TotalPrice: totalPrice,
				UnitPrice:  unitPrice,
				District:   community,
				SoldYear:   soldYear,
				SoldMonth:  soldMonth,
				Area:       area}
			err := db.Save(&sold).Error
			std.PrintErr(err)
		}
	})

	url := fmt.Sprintf("%s%s%s%d", sellUrl, community, "/pg", page)
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
}

func sellingInfo(community string, page int) {
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent(userAgent),
	)
	c.SetRequestTimeout(time.Duration(120) * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: sellUrl, Parallelism: 1}) //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
		llog.Info("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
	})
	//访问所有info 访问前20页采用goroutine
	c.OnHTML(".sellListContent>li", func(e *colly.HTMLElement) {
		re, _ := regexp.Compile(`\d+`)                                                                                                    //正则表达式用来匹配数字
		houseId := e.Attr("data-lj_action_housedel_id")                                                                                   //获取房子ID，可根据ID直接访问房子详情主页
		nameRegion := e.ChildText("div.info > div.flood > div.positionInfo > a")                                                          //同时获取小区名和详细地区
		name := strings.Split(nameRegion, " ")[0]                                                                                         //将同时获取的小区名和详细地区分离，取其中的小区名字
		region := strings.Split(nameRegion, " ")[1]                                                                                       //将同时获取的小区名和详细地区分离，取其中的详细地区
		totalPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .priceInfo .totalPrice span").Eq(0).Text()))))              //根据页面元素获取总价，正则匹配数字，转换成int类型
		unitPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .priceInfo .unitPrice span").Eq(0).Text()))))                //读取页面元素获取单价,正则匹配单价的数字，转换成int类型
		area, _ := strconv.Atoi(string(re.Find([]byte(strings.Split(e.ChildText("div.info > div.address > div.houseInfo "), " | ")[1])))) // //读取页面元素获取面积,正则匹配单价的数字，转换成int类型
		if houseId != "" {
			sell := Selling{
				Id:         houseId,
				Name:       name,
				TotalPrice: totalPrice,
				UnitPrice:  unitPrice,
				District:   community,
				Region:     region,
				Area:       area}
			err := db.Save(&sell).Error
			std.PrintErr(err)
		}
	})

	url := fmt.Sprintf("%s%s%s%d", sellUrl, community, "/pg", page)
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
}

//定义page结构体用来处理json
type Page struct {
	TotalPage int `json:"totalPage"`
	CurPage   int `json:"curPage"`
}

func sellingPage(community string) int {
	var totalPage int
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent(userAgent),
	)
	c.SetRequestTimeout(time.Duration(35) * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: sellUrl, Parallelism: 1}) //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
		llog.Info("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
	})
	//获取不同地区的总页数
	c.OnHTML(".contentBottom .house-lst-page-box", func(e *colly.HTMLElement) {
		page := Page{}
		err := json.Unmarshal([]byte(e.Attr("page-data")), &page)
		std.PrintErr(err)
		totalPage = page.TotalPage
	})

	url := fmt.Sprintf("%s%s", sellUrl, community)
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
	return totalPage
}

func soldPage(community string) int {
	var totalPage int
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent(userAgent),
	)
	c.SetRequestTimeout(time.Duration(90) * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: soldUrl, Parallelism: 1}) //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
		llog.Info("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
	})
	//获取不同地区的总页数
	c.OnHTML(".contentBottom .house-lst-page-box", func(e *colly.HTMLElement) {
		page := Page{}
		err := json.Unmarshal([]byte(e.Attr("page-data")), &page)
		std.PrintErr(err)
		totalPage = page.TotalPage
	})

	url := fmt.Sprintf("%s%s", soldUrl, community)
	c.OnError(func(_ *colly.Response, err error) {
		std.PrintErr(err)
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
	return totalPage
}
