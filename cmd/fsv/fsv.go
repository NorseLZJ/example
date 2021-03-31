package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

func showIp(port string) {
	addrS, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("Get InterfaceAddrs err:%v", err))
	}
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

type customFileHandler struct {
	root http.FileSystem
}

func CustomFileServer(root http.FileSystem) http.Handler {
	return &customFileHandler{root}
}

func (cf *customFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	ServeFile(w, r, cf.root, path.Clean(upath), true, "")
}

func ServeFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, name string, redirect bool, templateName string) {
	f, err := fs.Open(name)
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}

	if d.IsDir() {
		ListDirectory(w, r, f, templateName)
		return
	}
	http.ServeContent(w, r, d.Name(), d.ModTime(), f)
}

func toHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404 page not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}

func ListDirectory(w http.ResponseWriter, r *http.Request, f http.File, templateName string) {
	RootDir, err := f.Stat()
	if err != nil {
		panic(err)
	}
	var dirContents DirectoryContent
	dirContents.DirName = RootDir.Name()
	dirContents.Files = make([]FileContent, 0)
	dirs, err := f.Readdir(-1)
	if err != nil {
		log.Printf("http: error reading directory: %v", err)
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	for _, d := range dirs {
		name := d.Name()
		fileExtension := "page"
		if d.IsDir() {
			name += "/"
			fileExtension = "folder"
		} else if len(filepath.Ext(name)) > 1 {
			fileExtension = filepath.Ext(name)[1:]
		}

		url := url.URL{Path: name}
		fileContent := FileContent{Name: name, Size: GetHumanReadableSize(d), URL: url, Extension: fileExtension}
		dirContents.Files = append(dirContents.Files, fileContent)
	}
	dirContents.IPAddr = r.Host
	renderTemplate(w, templateName, dirContents)
}

type DirectoryContent struct {
	DirName string
	Files   []FileContent
	IPAddr  string
}

type FileContent struct {
	Name      string
	Size      string
	URL       url.URL
	Extension string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	var t *template.Template
	var err error
	// use default rendering html
	if len(tmpl) == 0 {
		t = template.New("index")
		t, err = t.Parse(DirListTemplateHTML)
	} else {
		templatePath, _ := filepath.Abs(tmpl + ".html")
		fmt.Println("template path", templatePath)
		t, err = template.ParseFiles(templatePath)
	}

	if err != nil {
		fmt.Println("Error in parsing template ", err)
		panic(err)
	}
	t.Execute(w, data)
}

func RequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func GetHumanReadableSize(f os.FileInfo) string {
	if f.IsDir() {
		return "--"
	}
	bytes := f.Size()
	mb := float32(bytes) / (1024.0 * 1024.0)
	return fmt.Sprintf("%.2f MB", mb)
}

var DirListTemplateHTML = `
<!DOCTYPE html><html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><title>Index of /</title><style>body{font:15px/1.4em Arial,"Helvetica Neue",Helvetica,sans-serif;padding:0 40px;background-color:#f2f2f2}.container{margin:80px auto 40px;max-width:960px;padding:40px 50px 30px;background-color:#fff}.badge{float:right}h1{font-size:28px;line-height:2em;margin:0}h2{font-size:18px;font-weight:400;margin:20px 0}.versions{float:right;padding:2px 12px 2px 6px;margin-top:14px;max-width:400px}.description{margin:10px 0;font-size:16px;color:#666}.path{margin:20px 0;padding:0;border-top:1px solid #e5e5e5;border-bottom:1px solid #e5e5e5}table{width:100%;border-spacing:0}.name{width:auto;text-align:left;padding-right:20px}.size{width:80px;text-align:right;padding-right:20px;color:#444}.time{width:240px;text-align:right;color:#444}th.name,th.size,th.time{color:#999;text-transform:uppercase;font-size:12px;letter-spacing:1px}a{color:#ff5627;text-decoration:none}a:focus,a:hover{color:#ff5627;text-decoration:underline}.landing{margin-top:30px;text-align:right}footer{max-width:960px;margin:0 auto 80px;font-size:14px;color:#666}.footer-left{float:left}.footer-right{float:right}.footer-right a{display:inline-block;margin-left:40px}</style></head><body><h1>Index of {{.DirName}}</h1><table>{{range .Files}}<tr><td class="icon-parent"><i class="fiv-cla fiv-icon-{{.Extension}}"></i></td><td class="file-size"><code>{{.Size}}</code></td><td class="display-name"><a href="{{.URL}}">{{.Name}}</a></td></tr>{{end}}</table><br><address style="font-size:1.5em"><a href="https://github.com/NorseLZJ/example/tree/master/cmd/fsv"><strong>fsv</strong></a> running @ {{.IPAddr}}</address></body></html>
`
