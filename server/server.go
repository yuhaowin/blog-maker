package server

import (
	"github.com/yuhaowin/blog-maker/render"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type ViewServer struct {
	ContentDir    string
	PostList      render.RenderList
	PostTemplate  *template.Template
	IndexTemplate *template.Template
}

// Viewing viewServer 结构体的方法
func (server *ViewServer) Viewing(writer http.ResponseWriter, request *http.Request) {
	var err error
	p := strings.TrimSpace(request.URL.Path)
	if p[len(p)-1] == '/' {
		err = render.GenerateListWithPath(server.IndexTemplate, server.PostList, p, writer)
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
