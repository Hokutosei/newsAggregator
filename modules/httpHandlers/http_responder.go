package httpHandlers

import (
	"encoding/json"
	"net/http"
)

// respondToJson write http json resposne
// Private func for package httpHandlers
func respondToJSON(w http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// PublicrespondToJSON write http json resposne
// Public func
func PublicrespondToJSON(w http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// JSONWriter write http json resposne
func JSONWriter(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		var s []byte
		return s
	}
	return json
}
