package httpHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web_apps/news_aggregator/modules/database"
)

type IndexVars struct {
	Ipaddress   string
	WebAppTitle string
	CurrentUser interface{}
}

var (
	db = database.MongodbSession
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("handled --> index")

	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	indexVars := IndexVars{"", "", nil}

	parsed_template_str := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsed_template_str)
	t.Execute(w, indexVars)
}
