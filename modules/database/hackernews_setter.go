package database

import(
	"fmt"
)

type HackerNews interface {}

var (
	hackerNewsCollection = "hackernews"
)

func HackerNewsInsert(hn HackerNews) {
	c := mongodbSession.DB(Db).C(hackerNewsCollection)
	err := c.Insert(hn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("saved!")
}
