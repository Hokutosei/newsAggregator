package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"web_apps/news_aggregator/modules/newsGetter"

	"web_apps/news_aggregator/modules/config"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/realtimeServer"
	_ "web_apps/news_aggregator/modules/utils"
)

var (
	serverPort = ":8000"
)

func handleAssets(assets ...string) {
	fmt.Println("called")
	for _, asset := range assets {
		assetDir := path.Join("public", asset)
		assetURLPath := fmt.Sprintf("/%s/", asset)
		//asset_dir := fmt.Sprintf("public/%s", asset)
		http.Handle(assetURLPath, http.StripPrefix(assetURLPath, http.FileServer(http.Dir(assetDir))))
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

	realtimeServer.Init()

	log.Println("now servering to port: ...", serverPort)
	http.ListenAndServe(serverPort, nil)
}
