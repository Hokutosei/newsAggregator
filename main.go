package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"web_apps/news_aggregator/modules/config"
	"web_apps/news_aggregator/modules/database"
	_ "web_apps/news_aggregator/modules/utils"
)

var (
	serverPort       = ":3000"
	loopCounterDelay = 300
)

// handleAssets serve all file assets
func handleAssets(assets ...string) {
	fmt.Println("called")
	for _, asset := range assets {
		start := time.Now()
		assetDir := path.Join("public", asset)
		assetURLPath := fmt.Sprintf("/%s/", asset)
		//asset_dir := fmt.Sprintf("public/%s", asset)
		http.Handle(assetURLPath, http.StripPrefix(assetURLPath, http.FileServer(http.Dir(assetDir))))
		fmt.Println(asset, " took: ", time.Since(start))
	}
}

// main entrypoint and main func for the app
func main() {
	fmt.Println("starting server....")

	go func() {
		fmt.Println("initializing backends...")
		config.StartEtcd()
		go database.MongodbStart()
		go database.StartRedis()

		// startRoutes start all routes
		startRoutes()

		assetsToHandle := []string{"images", "css", "js", "fonts"}
		handleAssets(assetsToHandle...)

		// news getter initializers
		// should set in admin page
		// go newsGetter.StartHackerNews(loopCounterDelay)
		// go newsGetter.StartGoogleNews(loopCounterDelay)
		InitNewRelic()
	}()

	fmt.Println("now servering to port -->>: ...", serverPort)
	http.ListenAndServe(serverPort, nil)
}
