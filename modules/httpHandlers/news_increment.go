package httpHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"web_apps/news_aggregator/modules/database"
)

// NewsIncrementParameter struct for news score incrementer
type NewsIncrementParameter struct {
	Id string
}

// IncrementNews incrase news score
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
