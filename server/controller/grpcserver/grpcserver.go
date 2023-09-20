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
	log   *zap.SugaredLogger
}

func NewGRPCServer(log *zap.SugaredLogger, auth *usecase.AuthUsecase, store *usecase.StorageUsecase, config *viper.Viper) *grpc.Server {
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
	grpcserver := grpc.NewServer()
	pb.RegisterMultiServiceServer(grpcserver, &grpcServer{log: log, auth: auth, store: store})
	return grpcserver
}

// RunGRPCServer запускает grpc
func RunGRPCServer(auth *usecase.AuthUsecase, store *usecase.StorageUsecase, log *zap.SugaredLogger, config *viper.Viper) (*grpc.Server, error) {
	grpcserver := NewGRPCServer(log, auth, store, config)
	addres := config.GetString("grpc_server.address")
	listen, err := net.Listen("tcp", addres)
	if err != nil {
		return nil, err
	}
	log.Infof("Running the server at: %s", addres)
	if err = grpcserver.Serve(listen); err != nil {
		return nil, err
	}
	return grpcserver, nil
}
