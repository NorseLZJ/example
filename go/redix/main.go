package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/golog"

	_ "github.com/davyxu/cellnet/peer/redix"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

var log = golog.New("server")

func main() {

	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("redix.Connector", "server", "127.0.0.1:6379", queue)

	//proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
	//	switch msg := ev.Message().(type) {
	//	case *cellnet.SessionAccepted:
	//		log.Debugln("server accepted")
	//	case *cellnet.SessionClosed:
	//		log.Debugln("session closed: ", ev.Session().ID())
	//	default:
	//		_ = msg
	//		fmt.Printf("unknow msg")
	//	}
	//})

	p.Start()
	queue.StartLoop()
	queue.Wait()
}
