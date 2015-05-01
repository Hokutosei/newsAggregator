package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"web_apps/news_aggregator/modules/news_getter"

	"web_apps/news_aggregator/modules/config"
	"web_apps/news_aggregator/modules/database"
	_ "web_apps/news_aggregator/modules/utils"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func handleAssets(assets ...string) {
	fmt.Println("called")
	for _, asset := range assets {
		asset_dir := path.Join("public", asset)
		asset_url_path := fmt.Sprintf("/%s/", asset)
		//asset_dir := fmt.Sprintf("public/%s", asset)
		http.Handle(asset_url_path, http.StripPrefix(asset_url_path, http.FileServer(http.Dir(asset_dir))))
	}
}

func main() {
	go func() {
		config.StartEtcd()
		go database.MongodbStart()

		startRoutes()

		assetsToHandle := []string{"images", "css", "js", "fonts"}
		handleAssets(assetsToHandle...)

		go news_getter.StartHackerNews()
		go news_getter.StartGoogleNews()

	}()

	log.Println("now servering to port 3000...")
	http.ListenAndServe(":3001", nil)
}
