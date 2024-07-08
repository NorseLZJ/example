package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	iris "github.com/kataras/iris/v12"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var (
	menu = flag.String("menu", "-", "start what framework [gin,iris,echo]")
	port = 3000

	client *rpc.Client = nil
)

func SendRpc() string {
	if client == nil {
		c, err := jsonrpc.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("Dialing:", err)
		}
		client = c
	}

	type RpcResult struct {
		Msg string
	}

retry:
	reply := &RpcResult{}
	if err := client.Call("GetInfo.Info", nil, &reply); err != nil {
		if errors.Is(err, rpc.ErrShutdown) {
			c, err := jsonrpc.Dial("tcp", "localhost:1234")
			if err != nil {
				log.Fatal("Dialing:", err)
			}
			client = c
			goto retry
		} else {
			return ""
		}
	}
	//fmt.Printf("GetInfo res: %v\n", reply)
	return reply.Msg
}

func startGin() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": SendRpc(),
		})
	})
	r.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

func startIris() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello <strong>%s</strong>!", "World")
	})

	app.Listen(fmt.Sprintf(":%d", port))
}

func startEcho() {
	e := echo.New()
	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func main() {
	flag.Parse()
	running := map[string]func(){
		"gin":  startGin,
		"echo": startEcho,
		"iris": startIris,
	}
	if rf, ok := running[*menu]; ok {
		rf()
	} else {
		flag.Usage()
	}
}
