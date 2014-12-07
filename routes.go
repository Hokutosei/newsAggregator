package main

import(
	"net/http"
	"fmt"
	"web_apps/news_aggregator/modules/http_handlers"
)

func startRoutes() {
	fmt.Println("load routes..")


	http.HandleFunc("/", http_handlers.Index)
}

//type IndexVars struct {
//	Ipaddress   string
//	WebAppTitle string
//	CurrentUser interface{}
//}
//
//func index(w http.ResponseWriter, r *http.Request) {
//	log.Println("handled --> index")
//
//	indexTemplate := "index.html"
//	t := template.New(indexTemplate).Delims("{{%", "%}}")
//	indexVars := IndexVars{"", "", nil}
//
//	parsed_template_str := fmt.Sprintf("public/%s", indexTemplate)
//	t, _ = t.ParseFiles(parsed_template_str)
//	t.Execute(w, indexVars)
//
//}
