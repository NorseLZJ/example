package main

import (
	"fmt"
	"time"
	"wscli/pb"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

var log = golog.New("client")

func sendMsg(callback func(string)) {
	for {
		time.Sleep(time.Second * 2)
		callback("hello world")
	}
}

func main() {

	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", "127.0.0.1:18801", queue)

	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			log.Debugln("client connected")
		case *cellnet.SessionClosed:
			log.Debugln("client error")
		case *pb.Ack:
			log.Infof("sid%d say: %s", msg.Id, msg.Msg)
		}
	})

	p.Start()
	queue.StartLoop()
	log.Debugln("Ready to chat!")
	sendMsg(func(str string) {
		ses := p.(interface{ Session() cellnet.Session }).Session()
		if ses == nil {
			fmt.Printf("ses is nil\n")
			return
		}
		req := &pb.Req{}
		req.Msg = str
		data, err := req.Marshal()
		if err != nil {
			fmt.Printf("Marshal err:%v\n", err)
			return
		}
		if ses != nil {
			ses.Send(req)
		}
		_ = data
	})
}
