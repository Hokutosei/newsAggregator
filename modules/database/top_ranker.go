package database

import (
	"fmt"
	"web_apps/news_aggregator/modules/cache"
)

var (
	todayTopRank = []string{"index", "news_top_rank"}
)

// TopRankingNews retrieve curated top ranking news fom cache
func TopRankingNews() (AggregatedNews, error) {
	fmt.Println("retreiving news cached")
	// RetrieveCachedNews(key string, redisPool *redis.Pool)
	key := cache.RedisKeyGen(todayTopRank...)

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
