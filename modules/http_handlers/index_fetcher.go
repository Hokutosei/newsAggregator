package http_handlers

import (
	_ "fmt"
	"net/http"
	_ "encoding/json"
	
	"gopkg.in/mgo.v2/bson"
)

type TestStruct struct {
	Status int
}

func indexNews() {

}

func GetIndexNews(w http.ResponseWriter, r *http.Request) {
	test := TestStruct{200}


	respondToJson(w, test)
}
