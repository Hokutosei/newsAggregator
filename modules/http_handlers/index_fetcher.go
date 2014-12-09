package http_handlers

import (
	_ "fmt"
	"net/http"
	_ "encoding/json"
	
	_ "gopkg.in/mgo.v2/bson"
	"web_apps/news_aggregator/modules/database"
)

type TestStruct struct {
	Status int
}

func indexNews() {

}

func GetIndexNews(w http.ResponseWriter, r *http.Request) {
	test := TestStruct{200}

	database.IndexNews()

	respondToJson(w, test)
}
