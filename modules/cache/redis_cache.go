package cache

import (
	"fmt"
	"sort"
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
	return convStrID(result...), nil
}

// RedisKeyGen is a util that joins slices to string
func RedisKeyGen(keys ...string) string {
	return strings.Join(keys, ":")
}

// convStrID
func convStrID(IDs ...string) []bson.ObjectId {
	fmt.Println(IDs)
	sort.Sort(sort.Reverse(sort.StringSlice(IDs)))
	fmt.Println(IDs)
	var objID []bson.ObjectId
	for _, i := range IDs {
		objID = append(objID, bson.ObjectIdHex(i))
	}
	return objID
}
