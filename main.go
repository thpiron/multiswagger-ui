package main

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/ovh/configstore"
)

//go:embed swagger.html
var htmlTmplt string

type OpenAPI struct {
	Name string `json:"name"`
	URL  string `json:"URL"`
}

type Conf struct {
	Title    string    `json:"title"`
	FavIcon  string    `json:"favicon"`
	OpenAPIs []OpenAPI `json:"openAPI"`
}

func main() {
	mux := http.NewServeMux()
	configstore.InitFromEnvironment()

	item, err := configstore.GetItemValue("config")
	if err != nil {
		log.Fatalf("error when retrieving config: %s", err)
		return
	}
	var conf Conf
	if err := json.Unmarshal([]byte(item), &conf); err != nil {
		log.Fatalf("error when unmarshaling the conf: %s", err)
		return
	}

	tmplt, err := template.New("swagger.html").Parse(htmlTmplt)

	if err != nil {
		log.Fatal("err when loading the template")
		return
	}

	mux.HandleFunc("/", handler(conf, tmplt))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")
}

func handler(conf Conf, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmplt.Execute(w, conf)
	}

}
