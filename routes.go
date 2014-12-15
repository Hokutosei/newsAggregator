package main

import (
	"fmt"
	"net/http"
	"web_apps/news_aggregator/modules/http_handlers"
)

func startRoutes() {
	fmt.Println("load routes..")

	http.HandleFunc("/", http_handlers.Index)
	http.HandleFunc("/get_index_news", http_handlers.LatestNews)
	http.HandleFunc("/top_score_news", http_handlers.GetIndexNews)
	http.HandleFunc("/latest_news", http_handlers.LatestNews)

	http.HandleFunc("/feed_more", http_handlers.FeedMore)
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
