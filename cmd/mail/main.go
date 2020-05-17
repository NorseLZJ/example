package main

import (
	"flag"
	"fmt"

	"github.com/NorseLZJ/example/std"
	"github.com/NorseLZJ/example/std/cfg_marshal"
	"github.com/astaxie/beego/utils"
)

const (
	cc = `{
"username":"%s",
"password":"%s",
"host":"%s",
"port":%d
}`
)

var (
	cf   = flag.String("f", "./mail.json", "cfg")
	cfgT = &cfg_marshal.SendMail{}
)

func main() {
	flag.Parse()
	err := cfg_marshal.Marshal(*cf, cfgT)
	std.CheckErr(err)
	sendMail()
}

func sendMail() {
	userName := cfgT.Your.User
	passWord := cfgT.Your.PassWord
	host := cfgT.Your.Host
	port := cfgT.Your.Port
	ccc := fmt.Sprintf(cc, userName, passWord, host, port)
	mail := utils.NewEMail(ccc)
	mail.To = make([]string, 0, len(cfgT.ToUserS))
	mail.To = append(mail.To, cfgT.ToUserS[:]...)
	mail.From = userName
	mail.Subject = cfgT.Title
	mail.Text = cfgT.Body
	//mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
	mail.AttachFile("")
	mail.Send()
}
