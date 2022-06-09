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

type viewingServer struct {
	ContentDir    string
	PostList      render.RenderList
	PostTemplate  *template.Template
	IndexTemplate *template.Template
}

var (
	addr          = ":8080"
	metaPath      = ".meta"
	draftFolder   = "draft"
	templatePath  = "templates"
	postTemplate  = "post.html.tpl"
	contentFolder = "content"
	indexTemplate = "index.html.tpl"
	outputFolder  = flag.String("o", "public", "")
)

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) == 2 {
		s := viewingServer{}
		s.PostList = make(render.RenderList)
		s.PostTemplate = render.GetTemplate(cwdPath(templatePath), postTemplate)
		s.IndexTemplate = render.GetTemplate(cwdPath(templatePath), indexTemplate)
		switch os.Args[1] {
		case "s":
			s.ContentDir = cwdPath(contentFolder)
			s.PostList.UpdateRenderList(s.ContentDir)
			log.Println("Starting HTTP server at http://localhost:8080/")
			log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(s.viewingServer)))
			return
		case "d":
			s.ContentDir = cwdPath(draftFolder)
			s.PostList.UpdateRenderList(s.ContentDir)
			log.Println("Starting HTTP server at http://localhost:8080/")
			log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(s.viewingServer)))
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

func (s *viewingServer) viewingServer(w http.ResponseWriter, r *http.Request) {
	var err error
	p := strings.TrimSpace(r.URL.Path)
	if p[len(p)-1] == '/' {
		err = render.GenerateListWithPath(s.IndexTemplate, s.PostList, p, w)
	} else if strings.Index(p, "/resource/images") == 0 {
		f, err := os.Open(cwdPath(p))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		io.Copy(w, f)
	} else {
		if p, ok := s.PostList[p]; ok {
			err = render.GeneratePostOut(s.PostTemplate, p.GetMDPath(s.ContentDir), w)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

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

func printHelp() {
	fmt.Println("./blog-maker d \n", "\trender all markdown files (in draft folder) and then start a HTTP server")
	fmt.Println("./blog-maker s \n", "\trender all markdown files and then start a HTTP server to exhibit your blog")
	fmt.Println("./blog-maker -o path \n", "\trender all markdown files (in content folder) to path, default path is ./public")
}
