package main

import (
	"flag"
	"fmt"
	"github.com/yuhaowin/blog-maker/render"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	addr          = ":8080"
	metaPath      = ".meta"
	templatePath  = "templates"
	postTemplate  = "post.html.tpl"
	contentFolder = "content"
	indexTemplate = "index.html.tpl"
	outputFolder  = flag.String("o", "public", "")
)

type viewingServer struct {
	ContentDir    string
	PostList      render.RenderList
	PostTemplate  *template.Template
	IndexTemplate *template.Template
}

// viewingServer 结构体的方法
func (server *viewingServer) viewing(writer http.ResponseWriter, request *http.Request) {
	var err error
	p := strings.TrimSpace(request.URL.Path)
	if p[len(p)-1] == '/' {
		err = render.GenerateListWithPath(server.IndexTemplate, server.PostList, p, writer)
	} else if strings.Index(p, "/resource/images") == 0 {
		f, err := os.Open(cwdPath(p))
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		io.Copy(writer, f)
	} else {
		if p, ok := server.PostList[p]; ok {
			err = render.GeneratePostOut(server.PostTemplate, p.GetMDPath(server.ContentDir), writer)
		} else {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	}
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

//----------------------------------------------------------------------------------------------------------------------

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) == 2 {
		server := viewingServer{}
		server.PostList = make(render.RenderList)
		server.PostTemplate = render.GetTemplate(cwdPath(templatePath), postTemplate)
		server.IndexTemplate = render.GetTemplate(cwdPath(templatePath), indexTemplate)
		switch os.Args[1] {
		case "s":
			server.ContentDir = cwdPath(contentFolder)
			server.PostList.UpdateRenderList(server.ContentDir)
			log.Println("Starting HTTP server at http://localhost:8080/")
			err := http.ListenAndServe(addr, http.HandlerFunc(server.viewing))
			log.Fatal(err)
			return
		case "h":
			printHelp()
			return
		default:
			printHelp()
			return
		}
	}
	flag.Parse()
	outputPath, _ := filepath.Abs(*outputFolder)
	render.Render(cwdPath(templatePath), cwdPath(contentFolder), filepath.Join(outputPath, metaPath), outputPath)
}

//有参数，有返回值
func cwdPath(subPath ...string) string {
	// using the function
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range subPath {
		workDir = filepath.Join(workDir, p)
	}
	return workDir
}

//无参数，无返回值
func printHelp() {
	fmt.Println("./blog-maker s \n", "\trender all markdown files and then start a HTTP server to exhibit your website")
	fmt.Println("./blog-maker -o path \n", "\trender all markdown files in content folder to path, default path is ./public")
}
