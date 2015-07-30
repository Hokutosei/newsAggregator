package httpHandlers

import (
	"net/http"
	"web_apps/news_aggregator/modules/database"
)

// Headlines news headlines handler
func Headlines(w http.ResponseWriter, r *http.Request) {
	headlines, err := database.HeadlinesGetter()
	if err != nil {
		return
	}

	respondToJSON(w, headlines)
}
