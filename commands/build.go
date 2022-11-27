package commands

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"time"
)

type PageData struct {
	Title      string
	SiteTitle  string
	Content    template.HTML
	PostTitles []string
}

type TitleWithTimestamp struct {
	Title     string
	Timestamp time.Time
}

var buildCmd = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {
		buildSite()
	},
}

func buildSite() {
	buildFrontPage()
	buildAboutPage()
	buildCvPage()
	buildProjectsPage()
	buildBlogPage()
}

func buildFrontPage() {
	f, err := os.Create(path.Join(BUILD_TARGET_PATH, "index.html"))

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	content := template.HTML(parseMarkdown(path.Join(CONTENT_PATH, "front_page.md")))

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		Content:   content,
	})

	if err != nil {
		log.Println("Error outputting rendered template!!")
	}
}

func buildAboutPage() {
	buildPath := path.Join(BUILD_TARGET_PATH, "about", "index.html")

	if _, err := os.Stat(path.Dir(buildPath)); err != nil {
		if err := os.Mkdir(path.Dir(buildPath), 0770); err != nil {
			log.Printf("Error creating directory: %s", buildPath)
		}
	}

	f, err := os.Create(buildPath)

	if err != nil {
		log.Println("Error building about page")
	}

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	content := template.HTML(parseMarkdown(path.Join(CONTENT_PATH, "about.md")))

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		Content:   content,
	})

}

func buildCvPage() {
	buildPath := path.Join(BUILD_TARGET_PATH, "cv", "index.html")

	if _, err := os.Stat(path.Dir(buildPath)); err != nil {
		if err := os.Mkdir(path.Dir(buildPath), 0770); err != nil {
			log.Printf("Error creating directory: %s", buildPath)
		}
	}

	f, err := os.Create(buildPath)

	if err != nil {
		log.Println("Error building cv page")
	}

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		Content:   template.HTML(parseMarkdown(path.Join(CONTENT_PATH, "cv.md"))),
	})
}

func buildProjectsPage() {
	buildPath := path.Join(BUILD_TARGET_PATH, "projects", "index.html")

	if _, err := os.Stat(path.Dir(buildPath)); err != nil {
		if err := os.Mkdir(path.Dir(buildPath), 0770); err != nil {
			log.Printf("Error creating directory: %s", buildPath)
		}
	}

	f, err := os.Create(buildPath)

	if err != nil {
		log.Println("Error building projects page")
	}

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		Content:   "All of my projects",
	})
}

func buildBlogPage() {

	buildPath := path.Join(BUILD_TARGET_PATH, "blog", "index.html")

	if _, err := os.Stat(path.Dir(buildPath)); err != nil {
		if err := os.Mkdir(path.Dir(buildPath), 0770); err != nil {
			log.Printf("Error creating directory: %s", buildPath)
		}
	}

	f, err := os.Create(buildPath)

	if err != nil {
		log.Println("Error building blog page")
	}

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	titlesByTimestamp := buildList(path.Join(CONTENT_PATH, "blog"))

	titles := make([]string, 0)
	for _, titleWithTimestamp := range titlesByTimestamp {
		titles = append(titles, titleWithTimestamp.Title)
	}

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:      "Welcome!",
		SiteTitle:  "Arnold Song",
		PostTitles: titles,
	})

}

func buildList(pathname string) []TitleWithTimestamp {

	files, err := ioutil.ReadDir(pathname)
	if err != nil {
		log.Printf("Error reading %s", pathname)
	}

	titlesByTimestamp := make([]TitleWithTimestamp, 0)
	for _, file := range files {
		metadata := parseMetadata(path.Join(pathname, file.Name()))
		postTitle := fmt.Sprintf("%s", metadata["Title"])
		timestamp, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", metadata["Timestamp"]))
		titlesByTimestamp = append(titlesByTimestamp,
			TitleWithTimestamp{
				Title:     postTitle,
				Timestamp: timestamp,
			},
		)
	}

	// Sort titles by time
	sort.Slice(titlesByTimestamp, func(i, j int) bool {
		return titlesByTimestamp[i].Timestamp.After(titlesByTimestamp[j].Timestamp)
	})

	return titlesByTimestamp
}

func parseMetadata(filename string) map[string]interface{} {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading %s", filename)
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err = markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		log.Println("Error converting to markdown")
	}
	return meta.Get(context)
}

func parseMarkdown(filename string) string {
	content, err := ioutil.ReadFile(filename)

	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.Emoji,
			extension.GFM,
			meta.Meta,
		),
	)

	if err != nil {
		log.Printf("Error reading: %s\n", filename)
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	if err = markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		log.Println("Error!!!")
	}
	metadata := meta.Get(context)
	if metadata != nil {
		log.Println(metadata["Title"])
	}

	return string(buf.Bytes())
}
