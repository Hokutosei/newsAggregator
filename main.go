package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"web_apps/news_aggregator/modules/config"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/security"
	"web_apps/news_aggregator/modules/utils"
)

var (
	serverPort       = ":3000"
	loopCounterDelay = 300
	hashKey          = "newsInstanceSecret"
	blockKey         = "newsInstanceBlock"
	cookieName       = "newsInstance.com"
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
		utils.Info(fmt.Sprintf("%s took: %v", asset, time.Since(start)))
	}
}

// main entrypoint and main func for the app
func main() {

	go func() {
		// config.StartConsul()
		config.StartEtcd()
		go database.MongodbStart()
		go database.StartRedis()

		// build secure cookies and keys
		security.BuildSecureKeys(hashKey, blockKey, cookieName)

		// startRoutes start all routes
		startRoutes()

		assetsToHandle := []string{"images", "css", "js", "fonts", "vendor"}
		handleAssets(assetsToHandle...)

		InitNewRelic()
	}()

	utils.Info(fmt.Sprintf("now servering to port -->> ... %v", serverPort))
	http.ListenAndServe(serverPort, nil)
}
