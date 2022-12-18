package interfaces

import (
	"context"
	"encoding/json"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/log"
	"github.com/ZuoFuhong/grpc-datacollector/server/errcode"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_datacollector_svr"
)

// ReportTrace 接收链路上报数据
func (s *GoDataCollectorImpl) ReportTrace(ctx context.Context, req *pb.ReportTraceReq) (*pb.ReportTraceRsp, error) {
	if req.GetTraceId() == "" {
		return nil, errcode.ErrLogicParam
	}
	reqBytes, _ := json.Marshal(req)
	if err := s.ts.AggregatedDocument(ctx, reqBytes); err != nil {
		log.Errorf("call AggregatedDocument failed, err: %v", err)
		return nil, errcode.ErrRpcRequestFail
	}
	return new(pb.ReportTraceRsp), nil
}
