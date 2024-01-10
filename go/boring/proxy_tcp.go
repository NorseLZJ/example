package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var (
	mu          sync.RWMutex
	portToMySQL = make(map[string]string)
)

func init() {
	// 初始化端口到 MySQL 地址的映射
	portToMySQL["3307"] = "mysql-server-ip1:3306"
	portToMySQL["3308"] = "mysql-server-ip2:3306"
	// 可以根据需要添加更多映射
}

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
	listenAddrs := []string{"0.0.0.0:3307", "0.0.0.0:3308"} // 代理监听地址列表

	for _, listenAddr := range listenAddrs {
		go func(addr string) {
			listener, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatal(err)
			}
			defer listener.Close()

			fmt.Printf("MySQL Proxy listening on %s...\n", addr)

			for {
				clientConn, err := listener.Accept()
				if err != nil {
					log.Println(err)
					continue
				}

				// 获取映射关系
				mu.RLock()
				targetAddr, ok := portToMySQL[addr]
				mu.RUnlock()

				if !ok {
					log.Printf("No MySQL address found for port %s\n", addr)
					clientConn.Close()
					continue
				}

				go handleClient(clientConn, targetAddr)
			}
		}(listenAddr)
	}

	// 等待程序退出
	select {}
}
