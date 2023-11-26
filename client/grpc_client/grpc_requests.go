package grpc_client

import (
	"context"
	"time"

	"github.com/gogapopp/gophkeeper/client/lib/random"
	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Hasher хеширует данные по секретному ключу
type Hasher interface {
	HashTextData(textdata models.TextData, userSecretPhrase string) (models.TextData, error)
	HashBinaryData(binarydata models.BinaryData, userSecretPhrase string) (models.BinaryData, error)
	HashCardData(carddata models.CardData, userSecretPhrase string) (models.CardData, error)
}

// Register отправляет запрос на регистрацию пользователя
func (g *GRPCClient) Register(ctx context.Context, user models.User) error {
	request := &pb.RegisterRequest{
		Login:      user.Login,
		Password:   user.Password,
		UserPhrase: user.UserPhrase,
	}
	_, err := g.grpcClient.Register(ctx, request)
	if err != nil {
		return err
	}
	return err
}

// Login  отправляет запрос на логин пользователя
func (g *GRPCClient) Login(ctx context.Context, user models.User) (*pb.LoginResponse, error) {
	request := &pb.LoginRequest{
		Login:      user.Login,
		Password:   user.Password,
		UserPhrase: user.UserPhrase,
	}
	response, err := g.grpcClient.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, err
}

// AddTextData отправляет запрос на сохранение текстовых данных на сервере
func (g *GRPCClient) AddTextData(ctx context.Context, textData models.TextData, userSecretPhrase string) error {
	textData, err := g.hashService.HashTextData(textData, userSecretPhrase)
	if err != nil {
		return err
	}
	textData.UploadedAt = timestamppb.New(time.Now())
	textData.UniqueKey = random.GenerateUniqueKey()
	request := &pb.TextDataRequest{
		TextData:   textData.TextData,
		UserID:     textData.UserID,
		Metainfo:   textData.Metainfo,
		UploadedAt: textData.UploadedAt,
		UniqueKey:  textData.UniqueKey,
	}
	_, err = g.grpcClient.AddTextData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddTextData(ctx, textData)
	if err != nil {
		return err
	}
	return err
}

// AddBinaryData отправляет запрос на сохранение бинарных данных на сервере
func (g *GRPCClient) AddBinaryData(ctx context.Context, binaryData models.BinaryData, userSecretPhrase string) error {
	binaryData, err := g.hashService.HashBinaryData(binaryData, userSecretPhrase)
	if err != nil {
		return err
	}
	binaryData.UploadedAt = timestamppb.New(time.Now())
	binaryData.UniqueKey = random.GenerateUniqueKey()
	request := &pb.BinaryDataRequest{
		BinaryData: binaryData.BinaryData,
		UserID:     binaryData.UserID,
		Metainfo:   binaryData.Metainfo,
		UploadedAt: binaryData.UploadedAt,
		UniqueKey:  binaryData.UniqueKey,
	}
	_, err = g.grpcClient.AddBinaryData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddBinaryData(ctx, binaryData)
	if err != nil {
		return err
	}
	return err
}

// AddCardData отправляет запрос на сохранение данных карты на сервер
func (g *GRPCClient) AddCardData(ctx context.Context, cardData models.CardData, userSecretPhrase string) error {
	cardData, err := g.hashService.HashCardData(cardData, userSecretPhrase)
	if err != nil {
		return err
	}
	cardData.UploadedAt = timestamppb.New(time.Now())
	cardData.UniqueKey = random.GenerateUniqueKey()
	request := &pb.CardDataRequest{
		CardNumberData: cardData.CardNumberData,
		CardNameData:   cardData.CardNameData,
		CardDateData:   cardData.CardDateData,
		CvvData:        cardData.CvvData,
		UserID:         cardData.UserID,
		Metainfo:       cardData.Metainfo,
		UploadedAt:     cardData.UploadedAt,
		UniqueKey:      cardData.UniqueKey,
	}
	_, err = g.grpcClient.AddCardData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddCardData(ctx, cardData)
	if err != nil {
		g.log.Info(err)
		return err
	}
	return err
}

// SyncData отправляет запрос на синхронизацию данных между клиентом и сервером по уникальным ключам пользователя
func (g *GRPCClient) SyncData(ctx context.Context, userID int) error {
	uniqueKeys, err := g.getService.GetUniqueKeys(ctx, userID)
	if err != nil {
		return err
	}
	uniqueKeysProto := make(map[string]*pb.RepeatedUniqueKeys)
	for key, values := range uniqueKeys {
		uniqueKeysProto[key] = &pb.RepeatedUniqueKeys{Values: values}
	}
	request := &pb.SyncRequest{
		Keys: uniqueKeysProto,
	}
	response, err := g.grpcClient.SyncData(ctx, request)
	if err != nil {
		g.log.Info(err)
		return err
	}
	err = g.saveService.SaveDatas(ctx, response)
	if err != nil {
		g.log.Info(err)
		return err
	}
	return err
}
