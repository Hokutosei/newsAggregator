package config

import (
	"fmt"
	"os"
	"web_apps/news_aggregator/modules/utils"
)

var (
	upstreamLimitCount = 10
	keyPrefix          = "newsaggregator_upstream_0"
	hostIPAddress      = ""
)

// RegisterServer to consul
func RegisterServer() {
	utils.Info("will register self..")

	utils.Info(fmt.Sprintf("%v local?", utils.IsLocal()))
	ScanCurrentUpstream()
}

// ScanCurrentUpstream scan current registered upstream if self is registered
func ScanCurrentUpstream() string {
	hostname := HostName()
	HostIPAddress(hostname)
	// newsaggregator_upstream_01
	// var key string
	for i := 0; i < upstreamLimitCount; i++ {
		ip, err := GetKV(fmt.Sprintf("%s%v", keyPrefix, i))
		if err != nil {
			break
		}
		utils.Info(fmt.Sprintf("val %s", ip))
	}

	utils.Info(hostname)

	return fmt.Sprintf("%v", keyPrefix)
}

// HostName get serverhostname
func HostName() string {
	hostname, _ := os.Hostname()
	return hostname
}

// HostIPAddress get host ipaddress
func HostIPAddress(hostname string) string {
	ip, err := GetKV(fmt.Sprintf("%s%s", hostIPAddress, hostname))
	if err != nil {
		return ""
	}
	utils.Info(fmt.Sprintf("host ip %s", ip))
	return ip
}
