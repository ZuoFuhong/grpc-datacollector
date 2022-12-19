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

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

// Serve 启动 Server
func (s *Server) Serve() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
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
		return err
	}

	esDb := es.NewESDb()
	serviceImpl := interfaces.InitializeService(esDb)
	gs := grpc.NewServer(grpc.UnaryInterceptor(gm.ChainUnaryServer(gr.UnaryServerInterceptor())))
	pb.RegisterGoDatacollectorSvrServer(gs, serviceImpl)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port))
	if err != nil {
		return err
	}
	glog.Debugf("Serving %s on %s", cfg.Server.Name, lis.Addr().String())
	return gs.Serve(lis)
}
