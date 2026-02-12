package render

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Render(tmplPath, contentPath, metaPath, outputPath string) {
	renderList := make(ContentList)
	err := readContentInfo(renderList, metaPath)
	if err != nil {
		log.Println(err.Error())
	}
	renderList.UpdateRenderList(contentPath)

	//check the output path does need to remove
	removedPosts := renderList.GetRemovedContentInfo(contentPath)
	for _, p := range removedPosts {
		removePath := filepath.Join(outputPath, p.IndexKey)
		log.Println("remove path:", removePath)
		if err = os.RemoveAll(removePath); err != nil {
			log.Println(err)
		}
	}

	doesNeedRenderAll := renderList.GetTemplateModifyTimes(tmplPath)
	if doesNeedRenderAll {
		log.Println("Render all files")
	}
	postTemplate := GetTemplate(tmplPath, "post.html.tpl")
	for key, info := range renderList {
		if info.IsContent() && (doesNeedRenderAll || info.NeedRender) {
			log.Printf("Rendering %s \n", info.GetMDOutPath(outputPath))
			err := GeneratePost(postTemplate,
				info.GetMDPath(contentPath),
				info.GetMDOutPath(outputPath))
			if err != nil {
				log.Fatal(err.Error())
				break
			}
			renderList[key].NeedRender = false
		}
	}
	log.Println("All md files has been rendered.")

	//generate homepage list, which sort by created time
	log.Println("Render index")
	sepratedList := make(map[string]ContentList)
	for k, v := range renderList {
		if !v.IsContent() {
			continue
		}
		paths := strings.Split(k, "/")
		var indexKey string
		// Handle different path structures:
		// /year/file -> ["", "year", "file"] -> len=3 -> indexKey="year"
		// /videos/year/file -> ["", "videos", "year", "file"] -> len=4 -> indexKey="videos/year"
		if len(paths) == 3 {
			// /year/file
			indexKey = paths[1]
		} else if len(paths) == 4 {
			// /category/year/file
			indexKey = paths[1] + "/" + paths[2]
		} else {
			indexKey = "/"
		}
		if _, ok := sepratedList[indexKey]; !ok {
			sepratedList[indexKey] = make(ContentList)
		}
		sepratedList[indexKey][k] = v
	}

	for k, v := range sepratedList {
		indexTemplate := GetTemplate(tmplPath, "index.html.tpl")
		err := GenerateList(indexTemplate, v, filepath.Join(outputPath, k, "index.html"))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Generate root index page with all blogs (excluding videos)
	log.Println("Render root index")
	allBlogsIndex := make(ContentList)
	for k, v := range sepratedList {
		if k != "videos" && !strings.Contains(k, "videos") {
			// Merge all year-based blogs
			for contentKey, contentVal := range v {
				allBlogsIndex[contentKey] = contentVal
			}
		}
	}
	if len(allBlogsIndex) > 0 {
		indexTemplate := GetTemplate(tmplPath, "index.html.tpl")
		err := GenerateList(indexTemplate, allBlogsIndex, filepath.Join(outputPath, "index.html"))
		if err != nil {
			log.Fatal(err)
		}
	}

	storeContentInfo(renderList, metaPath)
}
