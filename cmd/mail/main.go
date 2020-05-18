package main

import (
	"crypto/tls"
	"flag"

	"github.com/NorseLZJ/example/std"
	"github.com/NorseLZJ/example/std/cfg_marshal"
	"gopkg.in/gomail.v2"
)

var (
	cf  = flag.String("f", "./mail.json", "cfg")
	cft = &cfg_marshal.SendMail{}
)

func main() {
	flag.Parse()
	err := cfg_marshal.Marshal(*cf, cft)
	std.CheckErr(err)
	userName := cft.From
	passWord := cft.PassWord
	host := cft.Host
	port := cft.Port

	d := gomail.NewDialer(host, port, userName, passWord)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(genMail()...); err != nil {
		std.PrintErr(err)
	}
}

func genMail() []*gomail.Message {
	msgs := make([]*gomail.Message, 0, len(cft.To))
	title := cft.Title
	body := cft.Body
	file := cft.File
	from := cft.From
	for _, to := range cft.To {
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", to)
		//m.SetAddressHeader("", "dan@example.com", "Dan")
		m.SetHeader("Subject", title)
		m.SetBody("text/html", body)
		m.Attach(file)
		msgs = append(msgs, m)
	}
	return msgs
}
