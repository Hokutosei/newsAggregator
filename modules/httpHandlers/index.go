package httpHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/security"
	"web_apps/news_aggregator/modules/utils"
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
	utils.Info(fmt.Sprintf("main index page --> handled!"))
	// disable this for deployment, not finish yet
	// security.SetCookieHandler(w, r)
	security.ReadCookieHandler(w, r)

	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	indexVars := IndexVars{"", "", nil}

	parsedTemplateStr := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsedTemplateStr)
	t.Execute(w, indexVars)
}
