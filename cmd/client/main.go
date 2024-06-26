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

var (
	Version string
	Commit  string
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config, err := config.LoadConfig()
	fatal(err)
	log, err := logger.SetupLogger()
	fatal(err)
	repo, err := sqlite.NewRepo(config.GetString("grpc_client.clientDBdsn"))
	fatal(err)
	// закрываем подключение в БД при завершении программы
	defer func() {
		if err = repo.Close(); err != nil {
			fatal(err)
		}
	}()
	// получаем сервисы
	var (
		saveService = service.NewSaveService(repo)
		hashService = service.NewHashService()
		getService  = service.NewGetService(repo)
	)
	conn, err := grpc_client.ConnectGRPC(config)
	fatal(err)
	defer conn.Close()
	// подключаемся к серверу
	grpcclient, err := grpc_client.NewGRPCClient(conn, hashService, saveService, getService, log)
	fatal(err)
	// создаём приложение
	application := app.NewApplication(grpcclient, getService, Version, Commit, log)
	err = application.CreateApp()
	fatal(err)
	// реализация graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	log.Info("grpc client shutdown")
}
