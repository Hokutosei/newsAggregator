package http_handlers

import (
	"fmt"
	"net/http"
	_ "net/url"
	"encoding/json"
	
	_ "gopkg.in/mgo.v2/bson"
	"web_apps/news_aggregator/modules/database"
)

type TestStruct struct {
	Status int
}

type FeedMoreParams struct {
	ContentType string
	Skip	int
}

func indexNews() {

}

func GetIndexNews(w http.ResponseWriter, r *http.Request) {
	aggregated_news, err := database.NewsMainIndexNews()
	//aggregated_news, err := database.HackerNewsIndexNews()
	//aggregated_gn, err := database.GoogleNewsIndexNews()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondToJson(w, aggregated_news)
}

func LatestNews(w http.ResponseWriter, r *http.Request) {
	//aggregated_news, err := database.HackerNewsLatestNews()
	aggregated_news, err := database.NewsMainIndexNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondToJson(w, aggregated_news)
}

func FeedMore(w http.ResponseWriter, r *http.Request) {
	var feedMore FeedMoreParams
	if err := json.NewDecoder(r.Body).Decode(&feedMore); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	aggregated_news, err := database.HackerNewsFeedMore(feedMore.ContentType, feedMore.Skip)
	_ = err

	respondToJson(w, aggregated_news)

}

// func TopScoreNews(w http.ResponseWriter, r *http.Request) {

// }
