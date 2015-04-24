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
	session, err := mongodb.Dial("104.155.227.195:27020")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected to mongodb...")
	MongodbSession = session
}
