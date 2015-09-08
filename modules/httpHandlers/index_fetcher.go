package httpHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/models"
	"web_apps/news_aggregator/modules/utils"
)

// TestStruct a test Struct for utilities: to DEPRECATE
type TestStruct struct {
	Status int
}

// FeedMoreParams struct for feed more request
type FeedMoreParams struct {
	ContentType string
	Skip        int
}

// NewsIndexRequest struct for main news index request
type NewsIndexRequest struct {
}

// GetIndexNews get list down the index news
func GetIndexNews(w http.ResponseWriter, r *http.Request) {
	aggregatedNews, err := database.NewsMainIndexNews(defaultLang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// main json data output
	respondToJSON(w, aggregatedNews)
}

// LatestNews latest news getter to index page
// MAIN LATEST NEWS GETTER
func LatestNews(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	requestQuery := r.URL.Query()
	utils.Info(fmt.Sprintf("%v", requestQuery["lang"]))

	//insert language to query here
	aggregatedNews, err := database.NewsMainIndexNews(defaultLang)
	fmt.Println("FETCH index took: ", time.Since(start))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondToJSON(w, aggregatedNews)
}

// LatestNewsJSON retrieve latest news items in JSON only
func LatestNewsJSON(newsItems chan []byte) {
	start := time.Now()
	aggregatedNews, err := database.NewsMainIndexNews(defaultLang)
	fmt.Println("FETCH index took: ", time.Since(start))
	if err != nil {
		var x []byte
		newsItems <- x
		return
	}

	newsItems <- JSONWriter(aggregatedNews)
}

// TopScoreNews get news item that has greate news scores
func TopScoreNews(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	aggregatedNews, err := database.GetterNewsMainTopScore()
	fmt.Println("FETCH topScoreNews took: ", time.Since(start))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondToJSON(w, aggregatedNews)
}

// FeedMore fetch more news on offset
func FeedMore(w http.ResponseWriter, r *http.Request) {
	var feedMore FeedMoreParams
	if err := json.NewDecoder(r.Body).Decode(&feedMore); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	aggregatedNews, err := database.HackerNewsFeedMore(feedMore.ContentType, feedMore.Skip)
	_ = err

	respondToJSON(w, aggregatedNews)

}

// HeaderCategories list header topic categories
func HeaderCategories(w http.ResponseWriter, r *http.Request) {
	topics := models.TopicsList()

	respondToJSON(w, topics)
}

// FetchCategoryNews get categories news
func FetchCategoryNews(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	categoryInitial := r.URL.Query().Get("initial")

	categorizedNews, err := database.GetCategorizedNews(categoryInitial)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("fetchcategorynews TOOK: ", time.Since(start))
	respondToJSON(w, categorizedNews)
}
