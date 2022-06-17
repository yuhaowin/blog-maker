package render

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ContentInfo struct {
	NeedRender bool      //是否需要渲染
	Ext        string    //文件扩展名
	Title      string    //文章标题
	IndexKey   string    //排序字段
	CreateDate time.Time //创建时间
	ModifyTime time.Time //修改时间
}

func (c ContentInfo) IsContent() bool {
	return c.Ext == ".md"
}

func (c ContentInfo) GetMDPath(rootPath string) string {
	return filepath.Join(rootPath, c.IndexKey+c.Ext)
}

func (c ContentInfo) GetMDOutPath(rootPath string) string {
	return filepath.Join(rootPath, c.IndexKey, "index.html")
}

type ContentList map[string]*ContentInfo

func readContentInfo(list ContentList, path string) error {
	log.Println("Get all post update time from:", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewDecoder(file).Decode(&list)
}

func storeContentInfo(list ContentList, path string) error {
	log.Println("Save all post update time into:", path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		panic("open file error: " + err.Error())
		return err
	}
	defer file.Close()
	err = gob.NewEncoder(file).Encode(list)
	if err != nil {
		panic("open file error: " + err.Error())
		return err
	}
	return nil
}

func (list ContentList) GetTemplateModifyTimes(templatePath string) bool {
	needRenderALL := false
	filepath.Walk(templatePath, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}
		splits := strings.Split(path, templatePath)
		relativePath := splits[1]
		_, ok := list[relativePath]
		if !ok {
			list[relativePath] = &ContentInfo{
				IndexKey:   relativePath,
				ModifyTime: fileInfo.ModTime(),
			}
			needRenderALL = true
			return nil
		} else if !list[relativePath].ModifyTime.Equal(fileInfo.ModTime()) {
			list[relativePath].ModifyTime = fileInfo.ModTime()
			needRenderALL = true
			return nil
		}

		return nil
	})
	return needRenderALL
}

func (list ContentList) UpdateRenderList(contentPath string) {
	filepath.Walk(contentPath, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}

		splitIndex := strings.Index(path, contentPath)
		if splitIndex == -1 {
			log.Fatalf("split path error, content path=%s%t origin path=%s\n", contentPath, path)
			return nil
		}
		relativePath := path[len(contentPath):]
		fileExt := filepath.Ext(path)
		//remove the ext for key
		relativePath = relativePath[:len(relativePath)-len(fileExt)]

		if contentInfo, ok := list[relativePath]; ok {
			// do not need render
			if contentInfo.ModifyTime.Equal(fileInfo.ModTime()) {
				return nil
			}
		}
		if fileExt == ".md" {
			mdTitle, err := GetTitleFromPostMD(path)
			if err != nil {
				log.Fatal(err)
				return err
			}
			fileName := filepath.Base(path)
			dateSplits := strings.Split(fileName, "-")

			//parse create time
			d, err := time.Parse("2006-1-2", strings.Join(dateSplits[0:3], "-"))
			if err != nil {
				log.Fatalf("file %s must have create date, %s", fileName, err)
				return err
			}
			list[relativePath] = &ContentInfo{
				Title:      mdTitle,
				Ext:        fileExt,
				IndexKey:   relativePath,
				ModifyTime: fileInfo.ModTime(),
				CreateDate: d,
				NeedRender: true,
			}
		}
		return nil
	})
}

func (list ContentList) GetRemovedContentInfo(contentDir string) []*ContentInfo {
	var result []*ContentInfo
	for key, info := range list {
		if !info.IsContent() {
			continue
		}
		if _, err := os.Stat(info.GetMDPath(contentDir)); os.IsNotExist(err) {
			result = append(result, info)
			delete(list, key)
		}
	}
	return result
}
