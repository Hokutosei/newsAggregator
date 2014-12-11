package news_getter

import(
	"fmt"
	"encoding/json"
	"time"

	"web_apps/news_aggregator/modules/database"
)

var (
	loop_counter_delay = 5
	hacker_news_provider = "https://news.ycombinator.com"
	hacker_news_name	= "hackernews"
)

type HackerNewsTopStoriesId []int

type jsonNewsBody struct {
	By				string
	Id				int
	//Kids 			[]int
	Score			int
	Text			string
	Time			int
	Title			string
	Type			string
	Url				string
	ProviderName	string
	ProviderUrl		string
	CreatedAt		string
}

func StartHackerNews() {
	top_stories_ids := topStoriesId()
	content_out := make(chan jsonNewsBody)
	time_profiler := make(chan string)

	go func() {
		for t := range time.Tick(time.Duration(loop_counter_delay) * time.Second) {
			fmt.Println("running the loop")
			_ = t
			for _, id := range top_stories_ids {
				go func(id int, content_out chan jsonNewsBody, time_profiler chan string) {
					start := time.Now()
					news_content := hackerNewsReader(id)
					content_out <- news_content
					time_profiler <- fmt.Sprintf("%v", time.Since(start))
				}(id, content_out, time_profiler)
			}
		}
	}()

	go func() {
		for {
			content_out_msg := <- content_out
			time_profiler_out := <- time_profiler
			//fmt.Println(content_out_msg.Title)
			//fmt.Println(content_out_msg.Url)
			time_f := content_out_msg.Time
			content_out_msg.CreatedAt = fmt.Sprintf("%v", time.Now().Local())
			content_out_msg.ProviderUrl = hacker_news_provider
			content_out_msg.ProviderName = hacker_news_name
			//fmt.Println(time.Unix(int64(time_f), 0))
			_ = time_f

			// check if can save
			// then save
			can_save := database.HackerNewsFindIfExist(content_out_msg.Title)
			if can_save {
				database.HackerNewsInsert(content_out_msg)
			} else {
				//fmt.Println("did not save!")
			}
			_ = time_profiler_out
			//fmt.Println(time_profiler_out)
			//fmt.Println("----------------------------")
		}
	}()

}

func topStoriesId() []int {
	var top_stories_id_url string = "https://hacker-news.firebaseio.com/v0/topstories.json"
	var id_containers HackerNewsTopStoriesId
	response, _ := httpGet(top_stories_id_url)
	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &id_containers); err != nil {
		return id_containers
	}

	// make error handler
	return id_containers
}

func hackerNewsReader(id int) jsonNewsBody{
	news_url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	var news_content jsonNewsBody
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
