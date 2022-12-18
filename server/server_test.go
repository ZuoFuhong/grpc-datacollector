package server

import (
	"context"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_datacollector_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func Test_Client(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1218", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewGoDatacollectorSvrClient(conn)
	if _, err = client.ReportTrace(context.Background(), &pb.ReportTraceReq{
		TraceId:      "2ef9e9d0-82ff-4bbf-b205-5df5335aa5f1",
		Cmd:          "ImportWallet",
		Project:      "go_wallet_manage_svr",
		Source:       "grpc_gateway_best_practices",
		ServerIp:     "127.0.0.1",
		Errcode:      50000,
		Errmsg:       "系统请求错误，请稍后重试",
		Timestamp:    1671358211,
		Timecost:     150,
		Reqbody:      "{\"private_key\":\"0x01c4bda0939df07a31e3738c6c1e1d5905c9f229e6ffa1922557308a62efb23f\"}",
		SpanId:       2,
		PatentSpanId: 1,
	}); err != nil {
		t.Fatal(err)
	}
}
