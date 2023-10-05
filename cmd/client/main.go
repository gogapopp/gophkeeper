package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/gophkeeper/client/app"
	"github.com/gogapopp/gophkeeper/client/grpc_client"
	"github.com/gogapopp/gophkeeper/client/repository/sqlite"
	"github.com/gogapopp/gophkeeper/client/service"

	"github.com/gogapopp/gophkeeper/internal/config"
	"github.com/gogapopp/gophkeeper/internal/logger"
)

func main() {
	config, err := config.LoadConfig()
	fatal(err)
	log, err := logger.SetupLogger()
	fatal(err)
	repo, db, err := sqlite.NewRepo(config.GetString("grpc_client.clientDBdsn"))
	fatal(err)
	// закрываем подключение в БД при завершении программы
	defer func() {
		if err = db.Close(); err != nil {
			fatal(err)
		}
	}()
	saveService := service.NewSaveService(repo)
	hashService := service.NewHashService()
	getService := service.NewGetService(repo)
	conn, err := grpc_client.ConnectGRPC(config)
	fatal(err)
	defer conn.Close()
	grpcclient, err := grpc_client.NewGRPCClient(conn, hashService, saveService, getService, log)
	fatal(err)
	//
	// _ = grpcclient
	// uniqueKeys, err := getService.GetUniqueKeys(context.Background(), 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(uniqueKeys)
	// uniqueKeysProto := make(map[string]*pb.RepeatedUniqueKeys)
	// for key, values := range uniqueKeys {
	// 	uniqueKeysProto[key] = &pb.RepeatedUniqueKeys{Values: values}
	// }
	// request := &pb.SyncRequest{
	// 	Keys: uniqueKeysProto,
	// }
	// newclient := pb.NewMultiServiceClient(conn)
	// fmt.Println(request)
	// response, err := newclient.SyncData(context.Background(), request)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response, "123123123123123")
	// err = saveService.SaveDatas(context.Background(), response)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	application := app.NewApplication(grpcclient, getService, log)
	application.CreateApp()
	// реализация graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	log.Info("grpc client shutdown")
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
