package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Selling struct {
	Id         string    `gorm:"varchar(64) ;primary_key ;comment: '房子id'"`
	Name       string    `gorm:"varchar(64);comment:'小区名称'"`
	TotalPrice int       `gorm:"comment: '房子总价'"`
	UnitPrice  int       `gorm:"comment: '房子单价'"`
	District   string    `gorm:"varchar(64); comment:'所属行政区'"`
	Region     string    `gorm:"varchar(64); comment:'详细区域'"`
	Area       int       `gorm:"comment:'area'"`
	City       string    `gorm:"comment:'city'"`
	CreateTime time.Time `gorm:"comment:'createTime'"`
}

type Sold struct {
	Id         string    `gorm:"varchar(64) ;primary_key ;comment: '房子id'"`
	Name       string    `gorm:"varchar(64);comment:'小区名称'"`
	TotalPrice int       `gorm:"comment: '房子总价'"`
	UnitPrice  int       `gorm:"comment: '房子单价'"`
	District   string    `gorm:"varchar(64); comment:'所属行政区'"`
	SoldYear   string    `gorm:"varchar(32)  ;comment: '交易年份'"`
	SoldMonth  string    `gorm:"varchar(32)  ;comment: '交易月份'"`
	Area       int       `gorm:"comment:'面积'"`
	City       string    `gorm:"comment:'city'"`
	CreateTime time.Time `gorm:"comment:'createTime'"`
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
		fmt.Printf("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			log.Printf("OnError err:")
		}
	})
	//访问所有info 访问前20页采用goroutine
	c.OnHTML(".listContent>li", func(e *colly.HTMLElement) {
		re, _ := regexp.Compile(`\d+`)
		houseId := string(re.Find([]byte(strings.Split(e.ChildAttr("div.info > div.title > a", "href"), "/")[4])))
		name := strings.Split(e.ChildText("div.info > div.title > a"), " ")[0]
		area := 0
		if len(strings.Split(e.ChildText("div.info > div.title > a"), " ")) == 3 {
			area, _ = strconv.Atoi(string(re.Find([]byte(strings.Split(e.ChildText("div.info > div.title > a"), " ")[2]))))
		}
		totalPrice, _ := strconv.Atoi(e.DOM.Find(".info .address .totalPrice span").Eq(0).Text())
		unitPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .flood .unitPrice span").Eq(0).Text()))))
		dealDate := e.DOM.Find(".info .address .dealDate").Eq(0).Text()
		soldYear := strings.Split(dealDate, ".")[0]
		soldMonth := strings.Split(dealDate, ".")[1]
		if houseId != "" {
			val := &Sold{
				Id:         houseId,
				Name:       name,
				TotalPrice: totalPrice,
				UnitPrice:  unitPrice,
				District:   community,
				SoldYear:   soldYear,
				SoldMonth:  soldMonth,
				Area:       area,
				City:       city,
				CreateTime: createTime,
			}
			err := db.Save(val).Error
			if err != nil {
				fmt.Printf("err:%v", err)
			}
		}
	})

	url := fmt.Sprintf("%s%s%s%d", soldUrl, community, "/pg", page)
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err:%v", err)
		}
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
		fmt.Printf("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err:%v", err)
		}
	})
	c.OnHTML(".sellListContent>li", func(e *colly.HTMLElement) {
		re, _ := regexp.Compile(`\d+`)
		houseId := e.Attr("data-lj_action_housedel_id")
		nameRegion := e.ChildText("div.info > div.flood > div.positionInfo > a")
		name := strings.Split(nameRegion, " ")[0]
		region := strings.Split(nameRegion, " ")[1]
		totalPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .priceInfo .totalPrice span").Eq(0).Text()))))
		unitPrice, _ := strconv.Atoi(string(re.Find([]byte(e.DOM.Find(".info .priceInfo .unitPrice span").Eq(0).Text()))))
		area, _ := strconv.Atoi(string(re.Find([]byte(strings.Split(e.ChildText("div.info > div.address > div.houseInfo "), " | ")[1]))))
		if houseId != "" {
			val := &Selling{
				Id:         houseId,
				Name:       name,
				TotalPrice: totalPrice,
				UnitPrice:  unitPrice,
				District:   community,
				Region:     region,
				Area:       area,
				City:       city,
				CreateTime: createTime,
			}
			err := db.Save(val).Error
			if err != nil {
				fmt.Printf("err:%v", err)
			}
		}
	})

	url := fmt.Sprintf("%s%s%s%d", sellUrl, community, "/pg", page)
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err:%v", err)
		}
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
}

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
	c.Limit(&colly.LimitRule{DomainGlob: sellUrl, Parallelism: 1})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
	//获取不同地区的总页数
	c.OnHTML(".contentBottom .house-lst-page-box", func(e *colly.HTMLElement) {
		page := Page{}
		err := json.Unmarshal([]byte(e.Attr("page-data")), &page)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		totalPage = page.TotalPage
	})

	url := fmt.Sprintf("%s%s", sellUrl, community)
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err: %v", err)
		}
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
		fmt.Printf("Visiting :%s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
	//获取不同地区的总页数
	c.OnHTML(".contentBottom .house-lst-page-box", func(e *colly.HTMLElement) {
		page := Page{}
		err := json.Unmarshal([]byte(e.Attr("page-data")), &page)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		totalPage = page.TotalPage
	})

	url := fmt.Sprintf("%s%s", soldUrl, community)
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		c.Visit(url)
	})
	c.Visit(url)
	c.Wait()
	return totalPage
}
