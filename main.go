package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"web_apps/news_aggregator/modules/newsGetter"

	"web_apps/news_aggregator/modules/config"
	"web_apps/news_aggregator/modules/database"
	_ "web_apps/news_aggregator/modules/utils"
)

var (
	serverPort = ":3000"
)

// handleAssets serve all file assets
func handleAssets(assets ...string) {
	fmt.Println("called")
	for _, asset := range assets {
		assetDir := path.Join("public", asset)
		assetURLPath := fmt.Sprintf("/%s/", asset)
		//asset_dir := fmt.Sprintf("public/%s", asset)
		http.Handle(assetURLPath, http.StripPrefix(assetURLPath, http.FileServer(http.Dir(assetDir))))
	}
}

// main entrypoint and main func for the app
func main() {
	go func() {
		config.StartEtcd()
		go database.MongodbStart()

		// startRoutes start all routes
		startRoutes()

		assetsToHandle := []string{"images", "css", "js", "fonts"}
		handleAssets(assetsToHandle...)

		// news getter initializers
		// should set in admin page
		go newsGetter.StartHackerNews()
		go newsGetter.StartGoogleNews()

	}()

	InitNewRelic()

	log.Println("now servering to port: ...", serverPort)
	http.ListenAndServe(serverPort, nil)
}
