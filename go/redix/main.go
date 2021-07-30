package main

import (
	"time"

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

	conn, ok := p.(cellnet.RedisConnector)
	if !ok {
		panic("p is not redisConnector")
	}

	conn.SetConnectionCount(1)
	conn.SetDBIndex(2)
	p.Start()

	opt, ok := p.(cellnet.RedisPoolOperator)
	if !ok {
		panic("opt is not RedisPoolOperator")
	}
	SetClient(opt)

	time.AfterFunc(time.Second*2, func() {
		rcmd.Set("a", "b")
		rcmd.Set("c", "d")
		rcmd.Set("e", "f")
		rcmd.MSet("x", "1", "y", "2", "z", "3")
	})

	queue.StartLoop()
	queue.Wait()
}
