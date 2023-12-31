package grpc_client

import (
	"github.com/gogapopp/gophkeeper/client/service"
	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

// структура grpc клиента
type GRPCClient struct {
	grpcClient  pb.MultiServiceClient
	hashService *service.HashService
	saveService *service.SaveService
	getService  *service.GetService
	log         *zap.SugaredLogger
}

// ConnectGRPC пытается установить соединение с grpc сервером
func ConnectGRPC(config *viper.Viper) (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile("../../cert/server.crt", "")
	if err != nil {
		return nil, err
	}
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(config.GetString("grpc_client.address"), grpc.WithTransportCredentials(creds), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*1024)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewGRPCClient получаем клиент grpc сервера
func NewGRPCClient(conn *grpc.ClientConn, hashService *service.HashService, saveService *service.SaveService, getService *service.GetService, log *zap.SugaredLogger) (*GRPCClient, error) {
	newclient := pb.NewMultiServiceClient(conn)
	return &GRPCClient{
		grpcClient:  newclient,
		hashService: hashService,
		getService:  getService,
		saveService: saveService,
		log:         log,
	}, nil
}
