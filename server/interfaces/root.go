package interfaces

import (
	esinfra "github.com/ZuoFuhong/grpc-datacollector/server/infra/es"
	"github.com/ZuoFuhong/grpc-datacollector/server/service"
	"github.com/elastic/go-elasticsearch/v7"
)

type GoDataCollectorImpl struct {
	ts service.ITraceService
}

func NewDataCollectorImpl(ts service.ITraceService) *GoDataCollectorImpl {
	return &GoDataCollectorImpl{
		ts: ts,
	}
}

func InitializeService(esDb *elasticsearch.Client) *GoDataCollectorImpl {
	esInfra := esinfra.NewTraceIndex(esDb)
	ts := service.NewTraceService(esInfra)
	return NewDataCollectorImpl(ts)
}
