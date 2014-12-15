package news_getter

import(
	"fmt"
	"encoding/json"
	"time"
	"web_apps/news_aggregator/modules/database"
)

// type jsonNewsBody struct {
// 	By				string
// 	Id				int
// 	//Kids 			[]int
// 	Score			int
// 	Text			string
// 	Time			int
// 	Title			string
// 	Type			string
// 	Url				string
// 	ProviderName	string
// 	ProviderUrl		string
// 	CreatedAt		string
// }

type GoogleNews map[string]interface {}

type GoogleNewsResponseData struct {
	ResponseData struct {
		Results []GoogleNewsResults
	}
}

type ResponseData struct {
	Results []GoogleNewsResults
}	

type GoogleNewsResults struct {
	GsearchResultClass 	string
	ClusterUrl			string
	Content 			string
	UnescapedUrl		string
	Url					string
	Title				string
	TitleNoFormatting	string
	Publish				string
	PublishedDate		string
	Language			string
	RelatedStories 		[]RelatedStories
}

type RelatedStories struct {
	Url string
	TitleNoFormatting string
}

var (
	google_loop_counter_delay = 300
	google_news_provider = "https://news.google.com/"
	google_news_name	= "GoogleNews"

)


func StartGoogleNews() {
	url := fmt.Sprintf("https://ajax.googleapis.com/ajax/services/search/news?v=1.0&topic=t&ned=jp&userip=192.168.0.1")
	output_chan := make(chan GoogleNewsResults)

	go func() {
		for t := range time.Tick(time.Duration(google_loop_counter_delay) * time.Second) {
			_ = t
			//news_counter = 0
			go GoogleNewsRequester(url, output_chan)
		}
	}()

	go func() {
		for {
			output := <- output_chan
			GoogleNewsDataSetter(output)
			//fmt.Println("-------------------------")
		}
	}()

}

func GoogleNewsRequester(url string, output_chan chan GoogleNewsResults) {
	var google_news GoogleNewsResponseData
	response, _ := httpGet(url)
	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &google_news); err != nil {
		//return id_containers
		fmt.Println(err)
	}

	GNResponse := google_news.ResponseData
	for _, gn := range GNResponse.Results {
		output_chan <- gn
	}

}



func GoogleNewsDataSetter(google_news GoogleNewsResults) {

	can_save := database.GoogleNewsFindIfExist(google_news.Title)

	jsonNews := &jsonNewsBody{
		Title			: google_news.Title,
		By				: "GoogleNews",
		Score			: 0,
		Time			: int(time.Now().Unix()),
		Url				: google_news.Url,
		ProviderName	: "GoogleNews",
		RelatedStories	: google_news.RelatedStories,
		CreatedAt		: fmt.Sprintf("%v", time.Now().Local()),
	}

	if can_save {
		saved := database.GoogleNewsInsert(jsonNews)
		if saved {
			fmt.Println("saved!!")
			return
		}
		fmt.Println("did not save!")
	}
}
