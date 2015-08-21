package config

import (
	"encoding/json"
	"fmt"
	"os"

	"web_apps/news_aggregator/modules/utils"

	"github.com/coreos/go-etcd/etcd"
)

var (
	osMachine string

	machines []string
)

// EtcdResponse struct data from etcd response
type EtcdResponse struct {
	Action string `json:"action"`
	Node   struct {
		CreatedIndex  float64 `json:"createdIndex"`
		Key           string  `json:"key"`
		ModifiedIndex float64 `json:"modifiedIndex"`
		Value         string  `json:"value"`
	} `json:"node"`
}

// getOsMachinePrivateIP retrieve os machine private IP
func getOsMachinePrivateIP() {
	var envVar = "127.0.0.1"

	osEnvStr := []string{"COREOS_PRIVATE_IPV4", "LOCALHOST_IP"}
	for _, env := range osEnvStr {
		e := os.Getenv(env)
		if e != "" {
			envVar = e
			break
		}
	}

	machines = append(machines, fmt.Sprintf("http://%v:4001", envVar))
	fmt.Println("machine IP: ", machines)
}

// StartEtcd beginning connection
func StartEtcd() {
	utils.Info(fmt.Sprintf("starting etcd.."))
	getOsMachinePrivateIP()
}

// EtcdRawGetValue get raw or unmarshalled value from etcd cluster
func EtcdRawGetValue(key string) (string, error) {
	client := etcd.NewClient(machines)
	defer client.Close()

	val, err := client.RawGet(key, true, true)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var data EtcdResponse
	if err := json.Unmarshal(val.Body, &data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return data.Node.Value, nil
}
