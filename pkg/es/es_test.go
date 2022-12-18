package es

import (
	"github.com/ZuoFuhong/grpc-datacollector/pkg/config"
	"github.com/agiledragon/gomonkey"
	"log"
	"testing"
)

// go test -gcflags=all=-l -v -cover -run=Test_ES
func Test_ES(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyFunc(config.GlobalConfig, func() *config.Config {
		c := &config.Config{}
		c.Es.Address = "http://127.0.0.1:9200"
		return c
	})
	esdb := NewESDb()
	res, err := esdb.Info()

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
