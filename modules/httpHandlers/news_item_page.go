package httpHandlers

import (
	"fmt"
	"net/http"

	"web_apps/news_aggregator/modules/database"
)

// NewsItemParams struct for news item request
type NewsItemParams struct {
	NewsID string `json:"news_id"`
}

// NewsItemPage http handler for news item page request
func NewsItemPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("newsitemPage handler!")

	newsID := r.URL.Query().Get("news_id")
	result, err := database.NewsItemPage(newsID)
	if err != nil {
		return
	}

	respondToJSON(w, result)
	//NewsItemPage
}
