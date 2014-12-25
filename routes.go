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
	http.HandleFunc("/top_score_news", http_handlers.TopScoreNews)
	http.HandleFunc("/latest_news", http_handlers.LatestNews)

	http.HandleFunc("/feed_more", http_handlers.FeedMore)

	http.HandleFunc("/increment_news", http_handlers.IncrementNews)
}
