package httpHandlers

import (
	"fmt"
	"net/http"
)

// TopRankingNews handles top ranking news
func TopRankingNews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled top ranking news")
}
