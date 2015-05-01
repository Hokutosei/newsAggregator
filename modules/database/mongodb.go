package database

import (
	"fmt"

	mongodb "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"

	"web_apps/news_aggregator/modules/config"
)

var (
	MongodbSession *mongodb.Session

	Db                       = "news_aggregator"
	mongodbClusterKey string = "mongodb_cluster1"
)

// GetMongodbCluster retrieve mongodb cluster node from etcd
func GetMongodbCluster(host chan string) {
	mongodbCluster, err := config.EtcdRawGetValue(mongodbClusterKey)
	if err != nil {
		panic(err)
	}

	host <- mongodbCluster
}

// MongodbStart start connecting to mongodb
func MongodbStart() {
	fmt.Println("starting mongodb..")

	mongodbCluster := make(chan string)
	go GetMongodbCluster(mongodbCluster)

	host := <-mongodbCluster
	session, err := mongodb.Dial(host)
	if err != nil {
		// fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println("connected to mongodb...")
	MongodbSession = session
}

// SessionCopy make copy of a mongodb session
func SessionCopy() *mongodb.Session {
	return MongodbSession.Copy()
}
