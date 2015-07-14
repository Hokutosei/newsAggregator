package database

import (
	"fmt"
	"time"

	mongodb "gopkg.in/mgo.v2"

	"web_apps/news_aggregator/modules/config"
)

var (
	// MongodbSession main mongodb cluster session
	MongodbSession *mongodb.Session

	// Db current app database
	Db                = "news_aggregator"
	mongodbClusterKey = "mongodb_cluster1"
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

	mongoDBDialInfo := &mongodb.DialInfo{
		Addrs:   []string{host},
		Timeout: 20 * time.Second,
	}

	session, err := mongodb.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	fmt.Println("connected to mongodb...")
	MongodbSession = session
	// MongodbSession.SetMode(mongodb.Monotonic, true)
}

// SessionCopy make copy of a mongodb session
func SessionCopy() *mongodb.Session {
	return MongodbSession.Copy()
}
