package database

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"web_apps/news_aggregator/modules/config"

	utils "github.com/Hokutosei/hokutoseiUtils"
	"github.com/garyburd/redigo/redis"
)

var (
	// RedisPool main redis pool connection
	RedisPool *redis.Pool

	redisServer   = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "", "")
	redisHostKey  = "redisHost"
)

// Rstring struct for single KV
type Rstring struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

// StartRedis start connecting to redis
func StartRedis() {
	utils.Info(fmt.Sprintf("connecting to redis"))
	redisHost := make(chan string)
	go GetRedisHost(redisHost)

	s := <-redisHost
	RedisPool = NewPool(s)
	utils.Info(fmt.Sprintf("connected to redis!"))
}

// NewPool create redis pool servers
func NewPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// GetRedisHost get redis host from etcd
func GetRedisHost(host chan string) {
	redisHost, err := config.EtcdRawGetValue(redisHostKey)
	if err != nil {
		panic(err)
	}

	host <- redisHost
}

// RedisKeyGen is a util that joins slices to string
func RedisKeyGen(keys ...string) string {
	return strings.Join(keys, ":")
}

// Set insert to redis single KEY Value
func (r *Rstring) Set() (reply interface{}, err error) {
	conn := RedisPool.Get()
	defer conn.Close()

	n, err := conn.Do("SET", r.Key, r.Value)
	if err != nil {
		var s interface{}
		return s, err
	}
	return n, nil
}

// Get retrieve single KEY Value
func (r *Rstring) Get() (reply interface{}, err error) {
	conn := RedisPool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", r.Key))
	if err != nil {
		var x interface{}
		return x, err
	}
	return s, nil
}
