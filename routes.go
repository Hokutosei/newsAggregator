package main

import (
	"fmt"
	"net/http"
	http_handlers "web_apps/news_aggregator/modules/httpHandlers"
)

func startRoutes() {
	fmt.Println("load routes..")

	http.HandleFunc("/", http_handlers.Index)
	http.HandleFunc("/get_index_news", http_handlers.LatestNews)
	http.HandleFunc("/top_score_news", http_handlers.TopScoreNews)
	http.HandleFunc("/latest_news", http_handlers.LatestNews)

	http.HandleFunc("/news_item", http_handlers.NewsItemPage)

	http.HandleFunc("/feed_more", http_handlers.FeedMore)
	http.HandleFunc("/fetch_category_news", http_handlers.FetchCategoryNews)

	http.HandleFunc("/header_categories", http_handlers.HeaderCategories)

	http.HandleFunc("/increment_news", http_handlers.IncrementNews)
}
