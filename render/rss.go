package render

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

type RSSItem struct {
	Title       string
	Link        string
	Content     string
	PubDate     time.Time
	ContentPath string
}

func GenerateRSSOut(contentList ContentList, contentPath, siteURL, siteTitle, siteDescription string, writer io.Writer) error {
	var items []RSSItem
	for _, content := range contentList {
		if content.IsContent() {
			items = append(items, RSSItem{
				Title:       content.Title,
				Link:        siteURL + content.IndexKey + "/",
				ContentPath: content.GetMDPath(contentPath),
				PubDate:     content.CreateDate,
			})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PubDate.After(items[j].PubDate)
	})

	if len(items) > 20 {
		items = items[:20]
	}

	var rssItems []Item
	for _, item := range items {
		mdFile, err := ioutil.ReadFile(item.ContentPath)
		if err != nil {
			log.Printf("Error reading file %s: %v", item.ContentPath, err)
			continue
		}

		var buf bytes.Buffer
		if err = markdown.Convert(mdFile, &buf); err != nil {
			log.Printf("Error converting markdown %s: %v", item.ContentPath, err)
			continue
		}

		rssItems = append(rssItems, Item{
			Title:       item.Title,
			Link:        item.Link,
			Description: buf.String(),
			PubDate:     item.PubDate.Format(time.RFC1123Z),
			GUID:        item.Link,
		})
	}

	rss := RSS{
		Version: "2.0",
		Channel: &Channel{
			Title:       siteTitle,
			Link:        siteURL,
			Description: siteDescription,
			Language:    "zh-cn",
			Items:       rssItems,
		},
	}

	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		return err
	}

	xmlHeader := []byte(xml.Header)
	fullOutput := append(xmlHeader, output...)

	_, err = writer.Write(fullOutput)
	return err
}

func GenerateRSS(contentList ContentList, contentPath, outputPath, siteURL, siteTitle, siteDescription string) error {
	log.Println("Generating RSS feed")

	rssPath := filepath.Join(outputPath, "feed.xml")
	file, err := os.Create(rssPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = GenerateRSSOut(contentList, contentPath, siteURL, siteTitle, siteDescription, file)
	if err != nil {
		return err
	}

	log.Printf("RSS feed generated at %s", rssPath)
	return nil
}
