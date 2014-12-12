package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type GoogleNews interface {}


var(
	googleNewsCollection = "googleNews"
)

func GoogleNewsInsert(hn GoogleNews) bool {
	c := MongodbSession.DB(Db).C(googleNewsCollection)
	err := c.Insert(hn)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
//	fmt.Println("saved!")
}
func GoogleNewsFindIfExist(title string) bool {
	c := MongodbSession.DB(Db).C(googleNewsCollection)
	var result map[string]interface{}
	c.Find(bson.M{"title": title}).One(&result)
	if result["title"] != nil {
		return false
	}
	return true
}
