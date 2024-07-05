package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	iris "github.com/kataras/iris/v12"
	"github.com/labstack/echo/v4"
)

var (
	menu = flag.String("menu", "-", "start what framework [gin,iris,echo]")
	port = 3000
)

func startGin() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
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
