package main

import (
	"flag"
	"fmt"
	"github.com/yuhaowin/blog-maker/render"
	"github.com/yuhaowin/blog-maker/server"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	addr          = ":8080"
	metaPath      = ".meta"
	contentFolder = "content"
	templatePath  = "templates"
	postTemplate  = "post.html.tpl"
	indexTemplate = "index.html.tpl"
	outputFolder  = flag.String("o", "public", "")
)

//----------------------------------------------------------------------------------------------------------------------

func main() {
	log.SetFlags(log.Lshortfile)
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "s":
			log.Println("Starting HTTP server at http://localhost:8080/")
			err := startServer()
			log.Fatal(err)
		case "h":
			printHelp()
		default:
			printHelp()
		}
	} else {
		log.Println("Starting generation blog file...")
		flag.Parse()
		outputPath, _ := filepath.Abs(*outputFolder)
		render.Render(cwdPath(templatePath), cwdPath(contentFolder), filepath.Join(outputPath, metaPath), outputPath)
	}
}

func startServer() error {
	server := server.ViewServer{}
	server.PostList = make(render.ContentList)
	server.PostTemplate = render.GetTemplate(cwdPath(templatePath), postTemplate)
	server.IndexTemplate = render.GetTemplate(cwdPath(templatePath), indexTemplate)
	server.ContentDir = cwdPath(contentFolder)
	server.PostList.UpdateRenderList(server.ContentDir)
	return http.ListenAndServe(addr, http.HandlerFunc(server.Viewing))
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
