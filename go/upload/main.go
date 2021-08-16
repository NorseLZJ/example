package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	config = flag.String("c", "./config.json", "upload config")
)

func main() {
	flag.Parse()
	readConfig()
	if cfg.SaveDir == "" {
		cfg.SaveDir = Cwd()
	}
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/upload", func(c *gin.Context) {
		// single file
		key := c.PostForm("key")
		if key == "" || key != cfg.Key {
			return
		}

		file, _ := c.FormFile("file")
		dst := fmt.Sprintf("%s/%s", cfg.SaveDir, file.Filename)
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded err:%s!\n", file.Filename, err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded success!\n", file.Filename))
	})
	router.Run(cfg.Port)
}

func Cwd() string {
	cwd, _ := os.Getwd()
	return cwd
}
