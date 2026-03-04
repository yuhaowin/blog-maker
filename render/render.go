package render

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Render(tmplPath, contentPath, metaPath, outputPath string) {
	RenderWithConfig(tmplPath, contentPath, metaPath, outputPath, "", "", "")
}

func RenderWithConfig(tmplPath, contentPath, metaPath, outputPath, siteURL, siteTitle, siteDescription string) {
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

	// Extract years from renderList early
	yearsMap := make(map[string]bool)
	videoYearsMap := make(map[string]bool)
	for k, v := range renderList {
		if !v.IsContent() {
			continue
		}
		paths := strings.Split(k, "/")
		if len(paths) == 3 {
			// /year/file
			yearsMap[paths[1]] = true
		} else if len(paths) == 4 && paths[1] == "videos" {
			// /videos/year/file
			videoYearsMap[paths[2]] = true
		}
	}
	var years []string
	for year := range yearsMap {
		years = append(years, year)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(years)))
	var videoYears []string
	for year := range videoYearsMap {
		videoYears = append(videoYears, year)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(videoYears)))

	postTemplate := GetTemplate(tmplPath, "post.html.tpl")
	for key, info := range renderList {
		if info.IsContent() && (doesNeedRenderAll || info.NeedRender) {
			log.Printf("Rendering %s \n", info.GetMDOutPath(outputPath))
			err := GeneratePost(postTemplate,
				info.GetMDPath(contentPath),
				info.GetMDOutPath(outputPath),
				years, videoYears)
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
		err := GenerateList(indexTemplate, v, years, videoYears, filepath.Join(outputPath, k, "index.html"))
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
		err := GenerateList(indexTemplate, allBlogsIndex, years, videoYears, filepath.Join(outputPath, "index.html"))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Generate all-videos index page
	log.Println("Render videos index")
	allVideosIndex := make(ContentList)
	for k, v := range sepratedList {
		if strings.Contains(k, "videos") {
			for contentKey, contentVal := range v {
				allVideosIndex[contentKey] = contentVal
			}
		}
	}
	if len(allVideosIndex) > 0 {
		indexTemplate := GetTemplate(tmplPath, "index.html.tpl")
		err := GenerateList(indexTemplate, allVideosIndex, years, videoYears, filepath.Join(outputPath, "videos", "index.html"))
		if err != nil {
			log.Fatal(err)
		}
	}

	storeContentInfo(renderList, metaPath)

	// Generate RSS feed if site URL is provided
	if siteURL != "" {
		err := GenerateRSS(renderList, contentPath, outputPath, siteURL, siteTitle, siteDescription)
		if err != nil {
			log.Printf("Error generating RSS feed: %v", err)
		}
	}
}
