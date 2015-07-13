package cache

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
	// _ "web_apps/news_aggregator/modules/database"
	"github.com/garyburd/redigo/redis"
)

var (
	newsIndexKeySlice = []string{"index", "ids"}
)

// IndexNewsIDS main index news ids
func IndexNewsIDS(redisPool *redis.Pool) ([]bson.ObjectId, error) {
	start := time.Now()
	fmt.Println("retrieving news index ids")
	conn := redisPool.Get()
	defer conn.Close()

	key := RedisKeyGen(newsIndexKeySlice...)
	result, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		var x []bson.ObjectId
		return x, err
	}
	fmt.Println("indexnewsids took: ", time.Since(start))
	reversed := ReverseSlice(result...)
	// fmt.Println(result)
	return convStrID(reversed...), nil
}

// RedisKeyGen is a util that joins slices to string
func RedisKeyGen(keys ...string) string {
	return strings.Join(keys, ":")
}

// convStrID
func convStrID(IDs ...string) []bson.ObjectId {
	var objID []bson.ObjectId
	for _, i := range IDs {
		objID = append(objID, bson.ObjectIdHex(i))
	}
	return objID
}

// ReverseSlice util to reverse slice
func ReverseSlice(s ...string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}