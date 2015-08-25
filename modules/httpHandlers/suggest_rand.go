package httpHandlers

import (
	"fmt"
	"net/http"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/utils"
)

// SuggestRand suggest rand data
func SuggestRand(w http.ResponseWriter, r *http.Request) {
	utils.Info(fmt.Sprintf("suggest rand handled!"))

	news, err := database.SuggestRand()
	if err != nil {
		return
	}

	respondToJSON(w, news)
}
