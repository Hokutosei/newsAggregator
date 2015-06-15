package newsGetter

import (
	"encoding/json"
	"fmt"
	"time"
	"web_apps/news_aggregator/modules/database"
)

// GoogleNews interface for google news
type GoogleNews map[string]interface{}

// GoogleNewsResponseData response struct
type GoogleNewsResponseData struct {
	ResponseData struct {
		Results []GoogleNewsResults
	}
}

// ResponseData response struct
type ResponseData struct {
	Results []GoogleNewsResults
}

var (
	googleLoopCounterDelay = 300
	googleNewsProvider     = "https://news.google.com/"
	googleNewsName         = "GoogleNews"
)

// TopicsList return a list of topics/categories
func TopicsList() Topics {
	topics := Topics{
		"society":       TopicIdentity{"y", "社会"},
		"international": TopicIdentity{"w", "国際"},
		"business":      TopicIdentity{"b", "ビジネス"},
		"politics":      TopicIdentity{"p", "政治"},
		"entertainment": TopicIdentity{"e", "エンタメ"},
		"sports":        TopicIdentity{"s", "スポーツ"},
		"technology":    TopicIdentity{"t", "テクノロジー"},
		"pickup":        TopicIdentity{"ir", "ピックアップ"},
	}

	return topics
}

// StartGoogleNews start collecting google news
func StartGoogleNews() {
	fmt.Println("startgoogle news launched!")
	outputchan := make(chan GoogleNewsResults)

	go func() {
		for t := range time.Tick(time.Duration(googleLoopCounterDelay) * time.Second) {
			_ = t
			//news_counter = 0

			for k, v := range TopicsList() {
				go func(k string, v TopicIdentity) {
					url := fmt.Sprintf("https://ajax.googleapis.com/ajax/services/search/news?v=1.0&topic=%s&ned=jp&userip=192.168.0.1", v.Initial)
					GoogleNewsRequester(url, v, outputchan)
				}(k, v)
			}
		}
	}()

	go func() {
		for {
			output := <-outputchan
			GoogleNewsDataSetter(output)
			//fmt.Println("-------------------------")
		}
	}()
}

// GoogleNewsRequester google news http getter
func GoogleNewsRequester(url string, topic TopicIdentity, outputChan chan GoogleNewsResults) {
	var googleNews GoogleNewsResponseData
	response, err := httpGet(url)
	if err != nil {
		return
	}

	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &googleNews); err != nil {
		//return id_containers
		fmt.Println(err)
	}

	GNResponse := googleNews.ResponseData
	for _, gn := range GNResponse.Results {
		// set news item category
		gn.Category = topic

		// push to upstream channel
		outputChan <- gn
	}
}

// GoogleNewsDataSetter builds and construct data for insertion
func GoogleNewsDataSetter(googleNews GoogleNewsResults) {
	canSave := database.GoogleNewsFindIfExist(googleNews.Title)

	jsonNews := &jsonNewsBody{
		Title:          googleNews.Title,
		By:             "GoogleNews",
		Score:          0,
		Time:           int(time.Now().Unix()),
		Url:            googleNews.URL,
		ProviderName:   googleNewsName,
		RelatedStories: googleNews.RelatedStories,
		CreatedAt:      fmt.Sprintf("%v", time.Now().Local()),
		Category:       googleNews.Category,
	}

	// check if data exists already, need refactoring though
	if canSave {
		saved := database.GoogleNewsInsert(jsonNews)
		if saved {
			fmt.Println("saved!! google news!")
			return
		}
		fmt.Println("did not save!")
	}
}
