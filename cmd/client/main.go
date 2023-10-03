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
	conn, err := grpc_client.ConnectGRPC(config)
	fatal(err)
	defer conn.Close()
	grpcclient, err := grpc_client.NewGRPCClient(conn, hashService, saveService, log)
	fatal(err)
	application := app.NewApplication(grpcclient, log)
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