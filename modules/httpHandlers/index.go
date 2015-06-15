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

// Index news page
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("handled --> index")

	query := r.URL.Query()
	fmt.Println(query)

	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	indexVars := IndexVars{"", "", nil}

	parsedTemplateStr := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsedTemplateStr)
	t.Execute(w, indexVars)
}
