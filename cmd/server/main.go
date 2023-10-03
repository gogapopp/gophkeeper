package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/gophkeeper/internal/config"
	"github.com/gogapopp/gophkeeper/internal/logger"
	"github.com/gogapopp/gophkeeper/server/controller/grpcserver"
	"github.com/gogapopp/gophkeeper/server/repository/postgres"
	"github.com/gogapopp/gophkeeper/server/usecase"
)

func main() {
	config, err := config.LoadConfig()
	fatal(err)
	log, err := logger.SetupLogger()
	fatal(err)
	repo, db, err := postgres.NewRepo(config.GetString("grpc_server.serverDBdsn"))
	fatal(err)
	// закрываем подключение в БД при завершении программы
	defer func() {
		if err = db.Close(); err != nil {
			fatal(err)
		}
	}()
	// создаём сервис аутентификации
	authusecase := usecase.NewAuthUsecase(repo)
	// создаём сервис хранения данных
	storeusecase := usecase.NewStorageUsecase(repo)
	// запускаем сервер
	go func() {
		GRPCserver, err := grpcserver.RunGRPCServer(authusecase, storeusecase, log, config)
		fatal(err)
		defer GRPCserver.GracefulStop()
	}()
	// реализация graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	log.Info("grpc server shutdown")
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}