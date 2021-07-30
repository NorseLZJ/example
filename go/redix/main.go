package main

import (
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/golog"
	"github.com/mediocregopher/radix.v2/redis"

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

	time.AfterFunc(time.Second*2, func() {
		go func() {
			for {
				time.Sleep(time.Second)
				opt, ok := p.(cellnet.RedisPoolOperator)
				if !ok {
					panic("opt is not RedisPoolOperator")
				}
				opt.Operate(func(c interface{}) (ret interface{}) {
					if c == nil {
						log.Infof("redis client is nil")
						return nil
					}
					exec, ok := c.(*redis.Client)
					if !ok {
						log.Infof("client is not redis.Client")
						return nil
					}
					exec.Cmd("SET", "a", "b")
					return nil
				})
			}
		}()
	})
	queue.StartLoop()
	queue.Wait()
}
