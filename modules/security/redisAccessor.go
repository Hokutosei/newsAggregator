package security

import (
	"fmt"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/utils"
)

// GetSessionID retrieve from redis if session ID is pressent
func GetSessionID(key string) bool {
	utils.Info(fmt.Sprintf("getsessionId to redis"))
	rString := &database.Rstring{
		database.RedisKeyGen([]string{"sn", "session", key}...),
		"",
	}
	val, err := rString.Get()
	if err != nil {
		// create new key
		return false
	}
	utils.Info(fmt.Sprintf("session found! %v", val))
	return true
}
