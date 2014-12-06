package database

import(
	"fmt"

	mongodb "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var(
	mongodbSession mongodb.Session
)


func mongodbStart() {
	session, err := mongodb.Dial("130.211.246.179:49181")
	if err != nil {
		fmt.Println(err)
		return
	}
	mongodbSession = session
}
