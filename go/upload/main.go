package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	port    = flag.String("p", ":6500", "upload listen port")
	saveDir = flag.String("d", "", "upload save dir")
)

func main() {
	cwd := Cwd()
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		sign := c.PostForm("sign")
		if sign == "" || sign != "123456" {
			return
		}

		dst := ""
		if *saveDir == "" {
			dst = fmt.Sprintf("%s/%s", cwd, file.Filename)
		} else {
			dst = fmt.Sprintf("%s/%s", *saveDir, file.Filename)
		}
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(*port)
}

func Cwd() string {
	cwd, _ := os.Getwd()
	return cwd
}
