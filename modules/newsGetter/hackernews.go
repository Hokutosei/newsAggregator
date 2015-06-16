package newsGetter

import (
	"encoding/json"
	"fmt"
	"time"

	"web_apps/news_aggregator/modules/database"
)

var (
	loopCounterDelay   = 300
	hackerNewsProvider = "https://news.ycombinator.com"
	hackerNewsName     = "HackerNews"
)

// HackerNewsTopStoriesID struct for hacker news ids results
type HackerNewsTopStoriesID []int

// StartHackerNews starting GET hackernews
func StartHackerNews() {
	fmt.Println("starthacker news launched!")
	contentOut := make(chan jsonNewsBody)
	timeProfiler := make(chan string)

	go func() {
		for t := range time.Tick(time.Duration(loopCounterDelay) * time.Second) {
			topStoriesIds, err := topStoriesID()
			if err != nil {
				fmt.Println("skipping, err from topStoriesId")
				continue
			}
			fmt.Println("running the loop: ", t)

			for _, id := range topStoriesIds {
				go func(id int, contentOut chan jsonNewsBody, timeProfiler chan string) {
					start := time.Now()
					newsContent := hackerNewsReader(id)
					contentOut <- newsContent
					timeProfiler <- fmt.Sprintf("HN loop took: %v", time.Since(start))
				}(id, contentOut, timeProfiler)
			}
		}
	}()

	for {
		contentOutMsg := <-contentOut
		timeProfilerOut := <-timeProfiler

		timeF := contentOutMsg.Time
		contentOutMsg.Time = int(time.Now().Unix())
		contentOutMsg.CreatedAt = fmt.Sprintf("%v", time.Now().Local())
		contentOutMsg.ProviderUrl = hackerNewsProvider
		contentOutMsg.ProviderName = hackerNewsName

		_ = timeF

		// check if can save
		// then save
		canSave := database.HackerNewsFindIfExist(contentOutMsg.Title)
		if canSave {
			database.HackerNewsInsert(contentOutMsg)
		} else {
			//fmt.Println("did not save!")
		}
		_ = timeProfilerOut
		// fmt.Println(time_profiler_out)
		// fmt.Println("----------------------------")
	}
}

// topStoriesId
func topStoriesID() ([]int, error) {
	var topStoriesIDURL = "https://hacker-news.firebaseio.com/v0/topstories.json"
	var idContainers HackerNewsTopStoriesID
	response, err := httpGet(topStoriesIDURL)
	if err != nil {
		var x []int
		return x, err
	}

	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &idContainers); err != nil {
		return idContainers, nil
	}
	fmt.Printf("got %v ids:", len(idContainers))

	// make error handler
	return idContainers, nil
}

func hackerNewsReader(id int) jsonNewsBody {
	newsURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	var newsContent jsonNewsBody
	response, err := httpGet(newsURL)
	if err != nil {
		fmt.Println(err)
		return newsContent
	}
	defer response.Body.Close()

	contents, _ := responseReader(response)
	if err := json.Unmarshal(contents, &newsContent); err != nil {
		return newsContent
	}
	return newsContent
}
