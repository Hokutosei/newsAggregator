package news_getter

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type unMarshalledContent map[string]interface{}


func httpGet(url_string string) (*http.Response, error) {
	response, err := http.Get(url_string)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func responseReader(response *http.Response) ([]byte, error){
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func unmarshalResponseContent(content []byte, dataContainer interface {}) (interface {}, error) {
	if err := json.Unmarshal(content, &dataContainer); err !=nil {
		fmt.Println(err)
		return nil, err
	}
	return dataContainer, nil
}
