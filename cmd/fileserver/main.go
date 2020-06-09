package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/NorseLZJ/example/std"
)

var (
	port   = flag.String("p", ":18000", "server port")
	share  = flag.String("s", "", "public dir")
	upload = flag.String("u", "", "upload dir")
)

func main() {
	flag.Parse()
	if *share == "" {
		log.Fatal("share dir is nil")
	}
	fmt.Println("start server")
	showIp(*port)
	//http.HandleFunc("/", uploadFileHandler)
	//http.Handle("/file", http.StripPrefix("/file", http.FileServer(http.Dir(cfgT.ShareDir))))
	http.ListenAndServe(*port, http.FileServer(http.Dir(*share)))
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>文件互传</title>
</head>
<body style="text-align: center;"> 
    <h1>文件互传</h1>
    <br>
    <br>
    <form action="UploadFile.ashx" method="post" enctype="multipart/form-data">
    <input type="file" name="fileUpload" />
    <input type="submit" name="上传文件">
    </form>
        <br>
    <br>
        <br>
    <br>
    <a href="/file">文件下载</a>
</body>
</html>
        `)
	if r.Method == "POST" {
		file, handler, err := r.FormFile("fileUpload") //name的字段
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("err : %v", err)
			return
		}
		newFile, err := os.Create(*upload + handler.Filename)
		if err != nil {
			log.Printf("err : %v", err)
			return
		}
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil {
			std.CheckErr(err)
		}
		fmt.Println("upload successfully:" + *upload + handler.Filename)
		w.Write([]byte("SUCCESS"))
	}
}

func showIp(port string) {
	addrS, err := net.InterfaceAddrs()
	std.CheckErr(err)
	addrTmpS := make([]string, 0)
	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			addrTmpS = append(addrTmpS, ipNet.IP.String())
		}
	}
	fmt.Println("try access this address please")
	for _, v := range addrTmpS {
		fmt.Println(fmt.Sprintf("%s%s", v, port))
	}
}
