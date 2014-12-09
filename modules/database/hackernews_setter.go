package database

import(
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type HackerNews interface {}

var (
	hackerNewsCollection = "hackernews"
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

func IndexNews() {
	fmt.Println("called")
}
