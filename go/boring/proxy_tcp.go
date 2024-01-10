package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func copyData(dst, src net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Println(err)
	}
}

func handleClient(clientConn net.Conn, targetAddr string) {
	// 连接到 MySQL 服务器
	serverConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer serverConn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go copyData(clientConn, serverConn, &wg)
	go copyData(serverConn, clientConn, &wg)

	// 等待两个协程完成
	wg.Wait()
}

func main() {
	listenAddr := "0.0.0.0:8080"        // 代理监听地址
	targetAddr := "172.17.241.105:6379" // MySQL 服务器地址

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
	}()

	fmt.Printf("MySQL Proxy listening on %s...\n", listenAddr)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClient(clientConn, targetAddr)
	}
}
