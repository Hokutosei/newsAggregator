package database

import (
	"fmt"
	"web_apps/news_aggregator/modules/utils"
)

var (
	userSessionKey = "userSessionNews"
)

// SetSessionKey main session key setter to redis
func SetSessionKey(encodedKey string) bool {
	conn := RedisPool.Get()
	defer conn.Close()

	keySlice := []string{userSessionKey, SubstringStrToKey(encodedKey, 4)}
	key := RedisKeyGen(keySlice...)

	resp, err := conn.Do("SET", key, encodedKey)
	if err != nil {
		utils.Error(fmt.Sprintf("%v", err))
		return false
	}
	utils.Info(fmt.Sprintf("%v", resp))
	return true
}

// RetrieveSessionKey get session key from redis
func RetrieveSessionKey(cookieKey string) {
	conn := RedisPool.Get()
	defer conn.Close()

	eqIndex := utils.GetIndexOfCharInStr(cookieKey, "=")
	keySlice := RedisKeyGen([]string{userSessionKey, cookieKey[eqIndex:]}...)

	resp, err := conn.Do("GET", keySlice)
	if err != nil {
		utils.Error(fmt.Sprintf("%v", err))
		return
	}
	utils.Info(resp.(string))
}

// SubstringStrToKey get a key from string by range
func SubstringStrToKey(str string, keyLength int) string {
	return str[0:keyLength]
}
