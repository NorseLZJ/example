package main

import (
	"fmt"

	"wsgate/pb"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

var log = golog.New("server")

func main() {

	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "server", "127.0.0.1:18801", queue)

	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
		case *cellnet.SessionClosed:
			log.Debugln("session closed: ", ev.Session().ID())
		case *pb.Req:
			fmt.Printf("recv client req\n")
			ack := pb.Ack{
				Msg: msg.Msg,
				Id:  int32(ev.Session().ID()),
			}
			data, err := ack.Marshal()
			if err != nil {
				return
			}
			p.(cellnet.SessionAccessor).VisitSession(func(ses cellnet.Session) bool {
				ses.Send(data)
				return true
			})
			_ = msg
		default:
			fmt.Printf("unknow msg")
		}
	})

	p.Start()
	queue.StartLoop()
	queue.Wait()
}
