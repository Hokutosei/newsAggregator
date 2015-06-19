package httpHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web_apps/news_aggregator/modules/database"
)

// IndexVars used to be struct for index
// TO DEPRECATE
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

	fmt.Println(r.Header)

	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	indexVars := IndexVars{"", "", nil}

	parsedTemplateStr := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsedTemplateStr)
	t.Execute(w, indexVars)
}
