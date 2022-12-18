package service

import (
	"context"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/config"
	"github.com/ZuoFuhong/grpc-datacollector/server/document"
	esinfra "github.com/ZuoFuhong/grpc-datacollector/server/infra/es"
)

type ITraceService interface {
	// AggregatedDocument 聚合链路数据
	AggregatedDocument(ctx context.Context, reqBytes []byte) error
}

type TraceService struct {
	esInfra *esinfra.TraceIndex
}

func NewTraceService(esInfra *esinfra.TraceIndex) ITraceService {
	return &TraceService{
		esInfra: esInfra,
	}
}

// AggregatedDocument 聚合链路数据
func (s *TraceService) AggregatedDocument(ctx context.Context, reqBytes []byte) error {
	cfg := config.GlobalConfig()
	traceDoc := document.ConvertTraceDoc(reqBytes)
	return s.esInfra.WriteDocument(ctx, cfg.Es.Index, traceDoc.Bytes())
}
