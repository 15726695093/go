package server

import (
	"context"
	v1 "geektime/api/customer/v1"
	"geektime/app/customer/internal/conf"
	"geektime/app/customer/internal/service"
	"geektime/pkg/appmanage"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func (g *GrpcServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		g.server.Stop()
	}()
	return g.server.Serve(g.listener)
}

func NewGrpcServer(service *service.CustomerService, config *conf.GrpcConf) appmanage.GrpcServer {
	server := new(GrpcServer)
	lis, err := net.Listen("tcp", config.Addr())
	server.listener = lis
	if err != nil {
		panic(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterCustomerServiceServer(grpcServer, service)
	server.server = grpcServer
	return server
}
