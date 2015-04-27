package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

// HackerNews interface for hn news item
type HackerNews interface{}

// AggregatedNews interface struct for AggregatedNews
type AggregatedNews []interface{}

var (
	hackerNewsCollection = "news_main"
	searchLimitItems     = 50
)

// HackerNewsInsert insert data to mongodb
func HackerNewsInsert(hn HackerNews) {
	sc := SessionCopy()
	c := sc.DB(Db).C(hackerNewsCollection)
	defer sc.Close()

	err := c.Insert(hn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("saved! hackernews!")
}

// HackerNewsFindIfExist check if data exists already before saving
func HackerNewsFindIfExist(title string) bool {
	sc := SessionCopy()
	c := sc.DB(Db).C(hackerNewsCollection)
	defer sc.Close()

	var result map[string]interface{}
	c.Find(bson.M{"title": title}).One(&result)
	if result["title"] != nil {
		return false
	}
	return true
}

// HackerNewsIndexNews index page data
func HackerNewsIndexNews() (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(hackerNewsCollection)
	defer sc.Close()

	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Sort("-score").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}

// HackerNewsLatestNews aggregate latest news
func HackerNewsLatestNews() (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(hackerNewsCollection)
	defer sc.Close()

	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Sort("-_id").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}

// HackerNewsFeedMore feed more data on index list
func HackerNewsFeedMore(content_type string, length int) (AggregatedNews, error) {
	q := map[string]string{"latest_news": "-_id", "top_score_news": "-score"}
	sc := SessionCopy()
	c := sc.DB(Db).C(hackerNewsCollection)
	defer sc.Close()

	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Skip(length).Sort(q[content_type]).Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}
