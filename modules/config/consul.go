package config

import (
	"fmt"
	"os"
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
	pair, _, err := kv.Get(key, nil)
	if pair == nil || err != nil {
		return "", err
	}

	return string(pair.Value), nil
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
