package database

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"web_apps/news_aggregator/modules/cache"
	"web_apps/news_aggregator/modules/models"
	"web_apps/news_aggregator/modules/utils"
)

var (

	// NewsMainCollection collection name
	NewsMainCollection               = "news_main"
	dayHours           time.Duration = 24
	hoursPerDayQuery                 = dayHours * 1
)

// NewsMainIndexNews responder for index news query
// NEWS MAIN GETTER
func NewsMainIndexNews(language string) ([]models.NewsMaster, error) {
	newsIDChan := make(chan []bson.ObjectId)
	go cache.IndexNewsIDS(RedisPool, newsIDChan)

	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	var newsMaster []models.NewsMaster
	ids := <-newsIDChan
	err := c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&newsMaster)

	if err != nil {
		utils.Error(fmt.Sprintf("error getting main news index %v", err))
		return newsMaster, err
	}

	categorizeNewsItems(newsMaster...)
	return newsMaster, nil
}

func categorizeNewsItems(newsItems ...models.NewsMaster) {
	for _, news := range newsItems {
		fmt.Println(news.Category)
	}
}

// NewsMainIndexNewsCached retrieve index news from cached ID
// TODO make this func a util!
func NewsMainIndexNewsCached(IDs ...bson.ObjectId) (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	var aggregatedNews AggregatedNews
	err := c.Find(bson.M{"_id": bson.M{"$in": IDs}}).Sort("-score").Limit(5).All(&aggregatedNews)
	if err != nil {
		fmt.Println(err)
		return aggregatedNews, err
	}
	return aggregatedNews, nil
}

//GetterNewsMainTopScore main top page news getter
func GetterNewsMainTopScore() (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	var aggregatedNews AggregatedNews
	err := c.Find(bson.M{"url": bson.M{"$ne": ""}}).Sort("-score").Limit(searchLimitItems).All(&aggregatedNews)

	if err != nil {
		fmt.Println(err)
		return aggregatedNews, err
	}
	return aggregatedNews, nil
}

//IncrementNewsScore increment news score
// increment news ite page view
func IncrementNewsScore(paramsID string) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	err := c.Update(bson.M{"_id": bson.ObjectIdHex(paramsID)},
		bson.M{"$inc": bson.M{"score": 1}, "$currentDate": bson.M{"lastModified": true}})

	utils.HandleError(err)
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

	IncrementNewsScore(paramsID)
	return newsItem, nil
}

// GetCategorizedNews will get news with category news initials
func GetCategorizedNews(initial string) (AggregatedNews, error) {
	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	var aggregatedNews AggregatedNews
	err := c.Find(bson.M{"category.initial": initial}).Sort("-_id").Limit(searchLimitItems).All(&aggregatedNews)

	if err != nil {
		fmt.Println(err)
		return aggregatedNews, err
	}
	return aggregatedNews, nil
}
