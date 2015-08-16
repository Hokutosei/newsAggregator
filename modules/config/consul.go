package config

import (
	"fmt"
	"os"
	"time"
	"web_apps/news_aggregator/modules/utils"

	consul "github.com/hashicorp/consul/api"
)

var (
	consulEnv = "CONSUL_MASTER"

	// MongodBKey master mongodb key
	MongodBKey = "mongodb_cluster1"
	kv         *consul.KV
)

// StartConsul app session start for consul
func StartConsul() {
	// Get a new client, with KV endpoints
	client, _ := consul.NewClient(getConsulMasterIP())
	kv = client.KV()

}

// GetKV get k/v pair
func GetKV(key string) (string, error) {
	// Lookup the pair
	utils.Info(fmt.Sprintf("get KV: %s", key))
	retry := true
	retryCount := 0
	var pair *consul.KVPair
	for retry && retryCount != 5 {
		p, _, err := kv.Get(key, nil)
		if err != nil || p == nil {
			utils.Info(fmt.Sprintf("%v", err))
			utils.Info(fmt.Sprintf("%v", p))
			utils.Info(fmt.Sprintf("get KV: %s", key))
			utils.Info("debug ----------")
			retry = true
			time.Sleep(2000 * time.Millisecond)
			retryCount++
		} else {
			pair = p
			retry = false
		}
	}

	if pair == nil {
		panic("err key not found!")
	}

	result := string(pair.Value)
	utils.Info(fmt.Sprintf("found!: %v", result))
	return result, nil
}

// PutValue put/save k/v to consul
func PutValue(key, value string) {
	// PUT a new KV pair
	p := &consul.KVPair{Key: "foo", Value: []byte("test")}
	resp, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	utils.Info(fmt.Sprintf("put success! %v", resp))
}

func getConsulMasterIP() *consul.Config {
	e := os.Getenv(consulEnv)
	if e == "" {
		panic("consul not found!")
	}
	serverIP := fmt.Sprintf("%s:9200", e)
	config := consul.DefaultConfig()
	config.Address = serverIP
	fmt.Println(serverIP)
	return config
}
