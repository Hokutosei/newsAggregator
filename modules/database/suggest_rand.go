package database

import (
	"fmt"
	"web_apps/news_aggregator/modules/cache"
)

var (
	suggestRandKey = []string{"suggest", "random"}
)

// SuggestRand retrieve curated top ranking news fom cache
func SuggestRand() (AggregatedNews, error) {
	fmt.Println("retreiving news cached")
	key := cache.RedisKeyGen(suggestRandKey...)

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
