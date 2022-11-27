package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type HandleFunction func(writer http.ResponseWriter, request *http.Request)

type apiHandler struct {
	handler HandleFunction
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if h.handler != nil {
		log.Println("Using handler")
		h.handler(w, r)
	} else {
		fmt.Fprintf(w, "Default handler")
	}
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		paths := []string{"", "api", "about", "cv", "projects", "blog"}
		for _, path := range paths {
			createHandler(path)
		}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		buildSite()

		log.Printf("Starting server at %s:%s", host, port)
		log.Fatal(http.ListenAndServe(host+":"+port, nil))
	},
}

func init() {
	serverCmd.Flags().String("host", "", "Host address.  Default is 0.0.0.0")
	serverCmd.Flags().String("port", "8000", "Server port")
}

func createHandler(path string) {
	var url string
	h := apiHandler{}

	switch path {
	case "":
		url = "/"
		h.handler = rootHandlerFunc
	case "api":
		url = "/" + path + "/"
		h.handler = apiHandlerFunc
	default:
		url = "/" + path + "/"
		h.handler = indexHandlerFunc
	}
	log.Printf("Creating api handler for %s", url)
	http.Handle(url, h)
}

func rootHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	switch r.Method {
	case http.MethodGet:
		filename := "index.html"
		indexPath := path.Join(BUILD_TARGET_PATH, filename)

		f, err := ioutil.ReadFile(indexPath)
		if err != nil {
			log.Printf("Error reading #{indexPath}")
		}

		fmt.Fprint(w, string(f))
	}
}

func apiHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch path.Base(r.URL.Path) {
	case "menu":
		log.Printf("Give me the menu!!!")
	}
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		filename := "index.html"
		indexPath := path.Join(BUILD_TARGET_PATH, r.URL.Path, filename)

		f, err := ioutil.ReadFile(indexPath)
		if err != nil {
			log.Println("Error reading #{indexPath}")
		}

		fmt.Fprint(w, string(f))
	}
}
