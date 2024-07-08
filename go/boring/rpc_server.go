package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type GetInfo struct{}

type Result struct {
	Msg string
}

func (t *GetInfo) Info(_ interface{}, reply *Result) error {
	reply.Msg = fmt.Sprintf("is this rpc server! idx:%d\n", rand.Int())
	return nil
}

var (
	getInfoServer = new(GetInfo)
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	rpc.Register(getInfoServer)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println("Error listening:", e)
		return
	}
	defer l.Close()
	fmt.Println("Serving RPC server on port 1234")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			return
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
