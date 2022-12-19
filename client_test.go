package main

import (
	"context"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_datacollector_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1218", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewGoDatacollectorSvrClient(conn)
	if _, err = client.ReportTrace(context.Background(), &pb.ReportTraceReq{
		TraceId:      "3ef9e9d0-12ff-7bbf-b205-5df5335aa511",
		Cmd:          "ImportWallet",
		Project:      "go_wallet_manage_svr",
		Source:       "grpc-gateway-sample",
		ServerIp:     "127.0.0.1",
		Errcode:      0,
		Errmsg:       "",
		Timestamp:    time.Now().Unix(),
		Timecost:     150,
		Reqbody:      "{\"private_key\":\"0x01c4bda0939df07a31e3738c6c1e1d5905c9f229e6ffa1922557308a62efb23f\"}",
		SpanId:       2,
		PatentSpanId: 1,
	}); err != nil {
		t.Fatal(err)
	}
}
