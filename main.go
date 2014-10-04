package main

import (
	"net/http"
	"path"
	"fmt"
	"log"
)

type Profile struct {
	Name string
	Hobbies []string
}

func handleAssets(assets ...string) {
	for _, asset := range assets {
		asset_dir := path.Join("public", asset)
		asset_url_path := fmt.Sprintf("/%s/", asset)
		//asset_dir := fmt.Sprintf("public/%s", asset)
		//utilities.Logger(asset, nil)
		http.Handle(asset_url_path, http.StripPrefix(asset_url_path, http.FileServer(http.Dir(asset_dir))))
	}
}

func main() {

	assetsToHandle := []string{"images", "css", "js", "fonts"}
	go handleAssets(assetsToHandle...)

	log.Println("now servering to port 3000...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
