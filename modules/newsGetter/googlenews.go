package news_getter

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

// GoogleNewsResults google news result struct
type GoogleNewsResults struct {
	GsearchResultClass string
	ClusterUrl         string
	Content            string
	UnescapedUrl       string
	Url                string
	Title              string
	TitleNoFormatting  string
	Publish            string
	PublishedDate      string
	Language           string
	RelatedStories     []RelatedStories
	Image              Image
}

type RelatedStories struct {
	Url               string
	TitleNoFormatting string
}

type Image struct {
	Publisher string `json:"publisher"`
	URL       string `json:"url"`
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
	url := fmt.Sprintf("https://ajax.googleapis.com/ajax/services/search/news?v=1.0&topic=t&ned=jp&userip=192.168.0.1")
	outputchan := make(chan GoogleNewsResults)

	go func() {
		for t := range time.Tick(time.Duration(googleLoopCounterDelay) * time.Second) {
			_ = t
			//news_counter = 0
			go GoogleNewsRequester(url, outputchan)
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
func GoogleNewsRequester(url string, outputChan chan GoogleNewsResults) {
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
		Url:            googleNews.Url,
		ProviderName:   googleNewsName,
		RelatedStories: googleNews.RelatedStories,
		CreatedAt:      fmt.Sprintf("%v", time.Now().Local()),
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
