package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	"github.com/NorseLZJ/example/std"
)

var (
	port = flag.String("p", "8900", "server port")
	help = flag.Bool("h", false, "help ")
)

const (
	shareWindows = "D:\\share\\"
	shareLinux   = "~/share/"
	windows      = `windows`
	linux        = `linux`
)

var cc = `
you need create share dir before run fsv.exe 

usage: fsv [-p]
Example:
1、fsv 
2、fsv -p 9000 

Default:
Windows: share dir 	"D:\share\"
Linux: share dir 	"~/share/"
`

func main() {
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("%s\n", cc)
	}

	if *help {
		flag.Usage()
		return
	}
	share := ""
	switch runtime.GOOS {
	case windows:
		share = shareWindows
	case linux:
		share = shareLinux
	default:
		log.Fatal("share dir is nil")
	}

	fmt.Println("start server")
	*port = fmt.Sprintf(":%s", *port)
	showIp(*port)
	//http.HandleFunc("/", uploadFileHandler)
	//http.Handle("/file", http.StripPrefix("/file", http.FileServer(http.Dir(cfgT.ShareDir))))
	err := http.ListenAndServe(*port, http.FileServer(http.Dir(share)))
	if err != nil {
		log.Fatal(err)
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

//func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, `
//<!DOCTYPE html>
//<html lang="en">
//<head>
//    <meta charset="UTF-8">
//    <title>文件互传</title>
//</head>
//<body style="text-align: center;">
//    <h1>文件互传</h1>
//    <br>
//    <br>
//    <form action="UploadFile.ashx" method="post" enctype="multipart/form-data">
//    <input type="file" name="fileUpload" />
//    <input type="submit" name="上传文件">
//    </form>
//        <br>
//    <br>
//        <br>
//    <br>
//    <a href="/file">文件下载</a>
//</body>
//</html>
//        `)
//	if r.Method == "POST" {
//		file, handler, err := r.FormFile("fileUpload") //name的字段
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		fileBytes, err := ioutil.ReadAll(file)
//		if err != nil {
//			log.Printf("err : %v", err)
//			return
//		}
//		newFile, err := os.Create(*upload + handler.Filename)
//		if err != nil {
//			log.Printf("err : %v", err)
//			return
//		}
//		defer newFile.Close()
//		if _, err := newFile.Write(fileBytes); err != nil {
//			std.CheckErr(err)
//		}
//		fmt.Println("upload successfully:" + *upload + handler.Filename)
//		w.Write([]byte("SUCCESS"))
//	}
//}
