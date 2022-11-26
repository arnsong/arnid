package commands

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type PageData struct {
	Title      string
	SiteTitle  string
	Content    template.HTML
	PostTitles []string
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
		log.Println("Error building blog page")
	}

	tmpl := template.New("frontPage")
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "index.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "header.html")))
	tmpl = template.Must(tmpl.ParseFiles(path.Join(TEMPLATES_PATH, "content.html")))

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		Content:   "My work experience!!",
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
		log.Println("Error building blog page")
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

	err = tmpl.ExecuteTemplate(f, "index", PageData{
		Title:     "Welcome!",
		SiteTitle: "Arnold Song",
		PostTitles: []string{
			"Post 1", "Post 2", "Post 3",
		},
	})

}

func parseMarkdown(filename string) string {
	content, err := ioutil.ReadFile(filename)

	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.Emoji,
			extension.GFM,
		),
	)

	if err != nil {
		fmt.Println("Error reading: %s\n", filename)
	}

	var buf bytes.Buffer
	if err = markdown.Convert(content, &buf); err != nil {
		fmt.Println("Error!!!")
	}
	return string(buf.Bytes())
}
