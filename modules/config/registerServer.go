package config

import (
	"fmt"
	"web_apps/news_aggregator/modules/utils"
)

var (
	upstreamLimitCount = 10
	keyPrefix          = "/newsaggregator/upstream/"
)

// RegisterServer to consul
func RegisterServer() {
	utils.Info("will register self..")

	utils.Info(fmt.Sprintf("%v local?", utils.IsLocal()))
	ScanCurrentUpstream()
}

// ScanCurrentUpstream scan current registered upstream if self is registered
func ScanCurrentUpstream() string {
	// /newsaggregator/upstream/01
	var key string
	for i := 0; i < upstreamLimitCount; i++ {
		ip, err := GetKV(fmt.Sprintf("%s%v", keyPrefix, i))
		if err != nil {
			break
		}
		utils.Info(fmt.Sprintf("val %s", ip))
	}
	if key == nil {
		key = fmt.Sprintf("%v", keyPrefix)
	}
}
