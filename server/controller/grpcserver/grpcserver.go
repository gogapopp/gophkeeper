package grpcserver

import (
	"net"

	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/gogapopp/gophkeeper/server/usecase"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedMultiServiceServer
	auth  *usecase.AuthUsecase
	store *usecase.StorageUsecase
	get   *usecase.GetUsecase
	log   *zap.SugaredLogger
}

// NewGRPCServer создаём grpc сервер
func NewGRPCServer(log *zap.SugaredLogger, auth *usecase.AuthUsecase, store *usecase.StorageUsecase, get *usecase.GetUsecase, config *viper.Viper) *grpc.Server {
	// получаем сертификат из директории cert
	// creds, err := credentials.NewServerTLSFromFile("../../cert/server.crt", "../../cert/server.key")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// opts := []grpc.ServerOption{
	// 	grpc.Creds(creds),
	// }
	// opts...
	// создаём экземпляр grpc сервера
	grpcserver := grpc.NewServer(
		grpc.MaxRecvMsgSize(50*1024*1024),
		grpc.MaxSendMsgSize(1024*1024*1024),
	)
	pb.RegisterMultiServiceServer(grpcserver, &grpcServer{log: log, auth: auth, store: store, get: get})
	return grpcserver
}

// RunGRPCServer запускает grpc сервер
func RunGRPCServer(auth *usecase.AuthUsecase, store *usecase.StorageUsecase, log *zap.SugaredLogger, get *usecase.GetUsecase, config *viper.Viper) (*grpc.Server, error) {
	grpcserver := NewGRPCServer(log, auth, store, get, config)
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
