package grpc_client

import (
	"github.com/gogapopp/gophkeeper/client/service"
	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	grpcClient  pb.MultiServiceClient
	hashService *service.HashService
	saveService *service.SaveService
	log         *zap.SugaredLogger
}

func ConnectGRPC(config *viper.Viper) (*grpc.ClientConn, error) {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(config.GetString("grpc_client.address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewGRPCClient(conn *grpc.ClientConn, hashService *service.HashService, saveService *service.SaveService, log *zap.SugaredLogger) (*GRPCClient, error) {
	newclient := pb.NewMultiServiceClient(conn)
	return &GRPCClient{
		grpcClient:  newclient,
		hashService: hashService,
		saveService: saveService,
		log:         log,
	}, nil
}
