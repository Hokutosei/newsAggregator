package database

import (
	"fmt"
	"web_apps/news_aggregator/modules/cache"
)

var (
	headLines = []string{"weekly", "headlines"}
)

// HeadlinesGetter retrive data from mongodb
func HeadlinesGetter() (AggregatedNews, error) {
	fmt.Println("retreiving news headline cached")
	// RetrieveCachedNews(key string, redisPool *redis.Pool)
	key := cache.RedisKeyGen(headLines...)

	objID, err := cache.RetrieveCachedNews(key, RedisPool)
	if err != nil {
		fmt.Println(err)
		var x AggregatedNews
		return x, err
	}

	results, err := NewsMainIndexNewsCached(objID...)
	if err != nil {
		fmt.Println(err)
		return results, err
	}

	return results, nil
}
