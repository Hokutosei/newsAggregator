package http_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "gopkg.in/mgo.v2/bson"
	"web_apps/news_aggregator/modules/database"

	_ "web_apps/news_aggregator/modules/news_getter"
)

type NewsIncrementParameter struct {
	Id string
}

func IncrementNews(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var parameter NewsIncrementParameter
	err := decoder.Decode(&parameter)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parameter.Id)
	database.IncrementNewsScore(parameter.Id)
}

//func test(rw http.ResponseWriter, req *http.Request) {
//	decoder := json.NewDecoder(req.Body)
//	var t test_struct
//	err := decoder.Decode(&t)
//	if err != nil {
//		panic()
//	}
//	log.Println(t.Test)
//}
