package server

import (
	"fmt"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/config"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/es"
	glog "github.com/ZuoFuhong/grpc-datacollector/pkg/log"
	"github.com/ZuoFuhong/grpc-datacollector/server/interfaces"
	"github.com/ZuoFuhong/grpc-naming-monica/registry"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_datacollector_svr"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"net"
)

// RunServe 启动 Server
func RunServe() {
	cfg, err := config.LoadConfig()
	if err != nil {
		glog.Fatal("load config fail: " + err.Error())
	}
	config.SetGlobalConfig(cfg)

	// 服务注册
	if err := registry.NewRegistry(&registry.Config{
		Token:       cfg.Monica.Token,
		Namespace:   cfg.Monica.Namespace,
		ServiceName: cfg.Monica.ServiceName,
		IP:          cfg.Server.Addr,
		Port:        cfg.Server.Port,
	}).Register(); err != nil {
		glog.Fatal(err)
	}

	esDb := es.NewESDb()
	serviceImpl := interfaces.InitializeService(esDb)
	s := grpc.NewServer(grpc.UnaryInterceptor(gm.ChainUnaryServer(gr.UnaryServerInterceptor())))
	pb.RegisterGoDatacollectorSvrServer(s, serviceImpl)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port))
	if err != nil {
		glog.Fatal(err)
	}
	glog.Debugf("Serving %s on %s", cfg.Server.Name, lis.Addr().String())
	err = s.Serve(lis)
	if err != nil {
		glog.Fatal(err)
	}
}
