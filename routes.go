package main

import(
	"net/http"
	"html/template"
	"path"
	"log"

)


func index(w http.ResponseWriter, r *http.Request) {
	log.Println("handled --> index")

	profile := Profile{"jeane", []string{"programming", "gaming"} }

	fp := path.Join("public", "index.html")
	tmp, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmp.Execute(w, profile); err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
