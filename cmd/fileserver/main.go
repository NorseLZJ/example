package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/NorseLZJ/example/cfg_marshal"
)

var (
	share = flag.String("f", "./share.json", "cfg")
	cfgT  = &cfg_marshal.FileSrc{}
)

func main() {
	flag.Parse()
	fmt.Println("start server")
	err := cfg_marshal.Marshal(*share, cfgT)
	checkErr(err)
	showIp(cfgT.Port)
	http.HandleFunc("/", uploadFileHandler)
	http.Handle("/file", http.StripPrefix("/file", http.FileServer(http.Dir(cfgT.ShareDir))))
	log.Fatal(http.ListenAndServe(cfgT.Port, nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
		newFile, err := os.Create(cfgT.UpLoadDir + handler.Filename)
		if err != nil {
			log.Printf("err : %v", err)
			return
		}
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil {
			checkErr(err)
			return
		}
		fmt.Println("upload successfully:" + cfgT.UpLoadDir + handler.Filename)
		w.Write([]byte("SUCCESS"))
	}
}

func showIp(port string) {
	addrS, err := net.InterfaceAddrs()
	checkErr(err)
	addrTmpS := make([]string, 0)
	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			addrTmpS = append(addrTmpS, ipNet.IP.String())
		}
	}
	fmt.Println("try access this address please")
	for _, v := range addrTmpS {
		fmt.Println(fmt.Sprintf("%s:%s", v, port))
	}
}
