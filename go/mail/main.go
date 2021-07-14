package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

type Config struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	From     string   `json:"from"`
	PassWord string   `json:"password"`
	To       []string `json:"to"`
	Body     string   `json:"body"`
	Title    string   `json:"title"`
	File     string   `json:"file"`
}

var (
	cf      = flag.String("f", "./mail.json", "cfg")
	cft     = &Config{}
	timeStr = time.Now().Format("2006-01-02")
)

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*cf)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, cft)
	if err != nil {
		log.Fatal(err)
	}
	userName := cft.From
	passWord := cft.PassWord
	host := cft.Host
	port := cft.Port

	d := gomail.NewDialer(host, port, userName, passWord)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	d.DialAndSend(genMail()...)
}

func genMail() []*gomail.Message {
	msgs := make([]*gomail.Message, 0, len(cft.To))
	body := fmt.Sprintf("%s<br><br> From : %s <br> Date : %s", cft.Body, cft.From, timeStr)
	for _, to := range cft.To {
		m := gomail.NewMessage()
		m.SetHeader("From", cft.From)
		m.SetHeader("To", to)
		//m.SetAddressHeader("", "dan@example.com", "Dan")
		m.SetHeader("Subject", cft.Title)
		m.SetBody("text/html", body)
		m.Attach(cft.File)
		msgs = append(msgs, m)
	}
	return msgs
}
