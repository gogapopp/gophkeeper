package grpcserver

import (
	"net"

	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/gogapopp/gophkeeper/server/usecase"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcServer struct {
	pb.UnimplementedMultiServiceServer
	auth  *usecase.AuthUsecase
	store *usecase.StorageUsecase
	get   *usecase.GetUsecase
	log   *zap.SugaredLogger
}

// NewGRPCServer создаём grpc сервер
func NewGRPCServer(log *zap.SugaredLogger, auth *usecase.AuthUsecase, store *usecase.StorageUsecase, get *usecase.GetUsecase, config *viper.Viper) (*grpc.Server, error) {
	// получаем сертификат из директории cert
	creds, err := credentials.NewServerTLSFromFile("../../cert/server.crt", "../../cert/server.key")
	if err != nil {
		return nil, err
	}
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.MaxRecvMsgSize(1024 * 1024 * 1024),
		grpc.MaxSendMsgSize(1024 * 1024 * 1024),
	}
	// создаём экземпляр grpc сервера
	grpcserver := grpc.NewServer(
		opts...,
	)
	pb.RegisterMultiServiceServer(grpcserver, &grpcServer{log: log, auth: auth, store: store, get: get})
	return grpcserver, nil
}

// RunGRPCServer запускает grpc сервер
func RunGRPCServer(auth *usecase.AuthUsecase, store *usecase.StorageUsecase, log *zap.SugaredLogger, get *usecase.GetUsecase, config *viper.Viper) (*grpc.Server, error) {
	grpcserver, err := NewGRPCServer(log, auth, store, get, config)
	if err != nil {
		return nil, err
	}
	address := config.GetString("grpc_server.address")
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	log.Infof("Running the server at: %s", address)
	if err = grpcserver.Serve(listen); err != nil {
		return nil, err
	}
	return grpcserver, nil
}
