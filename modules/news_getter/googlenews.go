package news_getter

import(
	"fmt"
	"encoding/json"
	_ "time"
)

type GoogleNews map[string]interface {}

type GoogleNewsResponseData struct {
	ResponseData struct {
		Results []GoogleNewsResults
	}
}

type GoogleNewsResults struct {
	GsearchResultClass string
	ClusterUrl	string
	Content string
	UnescapedUrl	string
	Url	string
	Title	string
	TitleNoFormatting	string
	Publish	string
	PublishedDate	string
	Language	string
}

var (
	//loop_counter_delay = 300
	google_news_provider = "https://news.ycombinator.com"
	google_news_name	= "googlenew"
)


func StartGoogleNews() {
	url := fmt.Sprintf("https://ajax.googleapis.com/ajax/services/search/news?v=1.0&topic=t&ned=jp&userip=192.168.0.1")
	var google_news GoogleNews
	response, _ := httpGet(url)
	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &google_news); err != nil {
		//return id_containers
		fmt.Println(err)
	}
	//r := GoogleNewsResponseData{google_news}
	//fmt.Println(r.Results)
	fmt.Println(google_news)
	//response_data := google_news.ResponseData
	//fmt.Println(response_data.Results)
//	for _, s:= range response_data.Results {
//		fmt.Println(s.Title)
//	}
	//fmt.Println(google_news["responseData"]["results"])


	fmt.Println("start google news")
}


