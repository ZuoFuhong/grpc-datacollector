package es

import (
	"github.com/ZuoFuhong/grpc-datacollector/pkg/config"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/log"
	"github.com/elastic/go-elasticsearch/v7"
)

// NewESDb 创建 ES 连接实例
func NewESDb() *elasticsearch.Client {
	cfg := config.GlobalConfig()
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.Es.Address},
	})
	if err != nil {
		log.Fatal("Error creating the client: %s", err)
	}
	return es
}
