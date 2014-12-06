package news_getter

import(
	"fmt"
	"encoding/json"
	"time"

	_ "web_apps/news_aggregator/modules/database"
)

type HackerNewsTopStoriesId []int

type jsonNewsBody struct {
	By		string
	Id		int
	//Kids 	[]int
	Score	int
	Text	string
	Time	int
	Title	string
	Type	string
	Url		string
	CreatedAt	string


}

func StartHackerNews() {
	top_stories := topStoriesId()
	content_out := make(chan jsonNewsBody)
	for _, id := range top_stories {
		go func(id int, content_out chan jsonNewsBody) {
			news_content := hackerNewsReader(id)
			content_out <- news_content
		}(id, content_out)
	}
	go func() {
		for {
			content_out_msg := <- content_out
			fmt.Println(content_out_msg.Title)
			time_f := content_out_msg.Time
			content_out_msg.CreatedAt = time.Now().Local()
			fmt.Println(time.Unix(int64(time_f), 0))
			fmt.Println(content_out_msg.Score)
			//database.HackerNewsInsert(content_out_msg)
			fmt.Println("----------------------------")
		}
	}()

}

func topStoriesId() []int {
	var top_stories_id_url string = "https://hacker-news.firebaseio.com/v0/topstories.json"
	var id_containers HackerNewsTopStoriesId
	response, _ := httpGet(top_stories_id_url)
	defer response.Body.Close()

	contents, _ := responseReader(response)
	//top_stories_id, _ := unmarshalResponseContent(contents, id_containers)
	if err := json.Unmarshal(contents, &id_containers); err != nil {
		return id_containers
	}

	// make error handler
	return id_containers
}

func hackerNewsReader(id int) jsonNewsBody{
	news_url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	var news_content jsonNewsBody
	//fmt.Println(news_url)
	response, err := httpGet(news_url)
	if err != nil {
		//fmt.Println(err)
		return news_content
	}
	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &news_content); err != nil {
		return news_content
	}
	return news_content
}
