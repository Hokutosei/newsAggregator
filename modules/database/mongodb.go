package database

import (
	"fmt"

	mongodb "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var (
	MongodbSession *mongodb.Session

	Db = "news_aggregator"
)

func MongodbStart() {
	fmt.Println("starting mongodb..")
	session, err := mongodb.Dial("130.211.246.179:49181")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected to mongodb...")
	MongodbSession = session
}
