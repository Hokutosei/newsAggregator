package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var (
	NewsMainCollection = "news_main"
)

func NewsMainIndexNews() (AggregatedNews, error) {
	c := MongodbSession.DB(Db).C(NewsMainCollection)
	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Sort("-_id").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}

func GetterNewsMainTopScore() (AggregatedNews, error) {
	c := MongodbSession.DB(Db).C(NewsMainCollection)
	var aggregated_news AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Sort("-score").Limit(searchLimitItems).All(&aggregated_news)

	if err != nil {
		fmt.Println(err)
		return aggregated_news, err
	}
	return aggregated_news, nil
}
