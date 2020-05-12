package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gpmgo/gopm/modules/log"

	"github.com/gocolly/colly"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func main() {
	//llog := log2.LLog()
	//resp, err := http.Get("http://www.zhenai.com/zhenghun") //获取页面返回的response
	url := "https://xa.lianjia.com/ershoufang/"
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	log.Fatal("resp.StatusCode :%d", resp.StatusCode)
	//}
	//utf8Reader := transform.NewReader(resp.Body, coder(resp.Body).NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}
	//fd, err := os.OpenFile("tmp.html", os.O_WRONLY|os.O_TRUNC, 0600)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("%s", all)
	//defer fd.Close()
	//fd.Write(all)

	// colly
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		log.Info("%s", e.Name)
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
}

func coder(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
