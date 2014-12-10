package database

import(
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type HackerNews interface {}

type AggregatedNews []interface {}

var (
	hackerNewsCollection = "hackernews"
	searchLimitItems	= 50
)

func HackerNewsInsert(hn HackerNews) {
	c := MongodbSession.DB(Db).C(hackerNewsCollection)
	err := c.Insert(hn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("saved!")
}
func HackerNewsFindIfExist(title string) bool {
	c := MongodbSession.DB(Db).C(hackerNewsCollection)
	var result map[string]interface{}
	c.Find(bson.M{"title": title}).One(&result)
	if result["title"] != nil {
		return false
	}
	return true
}

func HackerNewsIndexNews() (AggregatedNews, error){
	c := MongodbSession.DB(Db).C(hackerNewsCollection)
	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": "null"}}).Sort("-score").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}

func HackerNewsLatestNews() (AggregatedNews, error) {
	c := MongodbSession.DB(Db).C(hackerNewsCollection)
	var aggregated_news AggregatedNews
	err := c.Find(nil).Sort("-_id").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}

