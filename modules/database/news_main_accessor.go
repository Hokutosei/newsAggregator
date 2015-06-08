package database

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"web_apps/news_aggregator/modules/utils"
)

var (

	// NewsMainCollection collection name
	NewsMainCollection = "news_main"
)

// NewsMainIndexNews responder for index news query
func NewsMainIndexNews() (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

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

func IncrementNewsScore(params_id string) {
	c := MongodbSession.DB(Db).C(NewsMainCollection)
	var aggregated_news interface{}
	fmt.Println(params_id)

	err := c.Update(bson.M{"_id": bson.ObjectIdHex(params_id)},
		bson.M{"$inc": bson.M{"score": 1}, "$currentDate": bson.M{"lastModified": true}})

	utils.HandleError(err)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(params_id)}).One(&aggregated_news)
	utils.HandleError(err)

	fmt.Println(aggregated_news)
}

// NewsItemPage get news item data
func NewsItemPage(paramsID string) (interface{}, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	var newsItem interface{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(paramsID)}).One(&newsItem)

	if err != nil {
		return newsItem, err
	}

	return newsItem, nil
}
