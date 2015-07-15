package httpHandlers

import (
	"fmt"
	"net/http"
	"web_apps/news_aggregator/modules/database"
)

// TopRankingNews handles top ranking news
func TopRankingNews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled top ranking news")

	database.TopRankingNews()

}
