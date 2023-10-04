package grpc_client

import (
	"context"
	"time"

	"github.com/gogapopp/gophkeeper/client/lib/random"
	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Hasher interface {
	HashTextData(textdata models.TextData, userSecretPhrase string) (models.TextData, error)
	HashBinaryData(binarydata models.BinaryData, userSecretPhrase string) (models.BinaryData, error)
	HashCardData(carddata models.CardData, userSecretPhrase string) (models.CardData, error)
}

func (g *GRPCClient) Register(ctx context.Context, user models.User) error {
	request := &pb.RegisterRequest{
		Login:    user.Login,
		Password: user.Password,
	}
	_, err := g.grpcClient.Register(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCClient) Login(ctx context.Context, user models.User) (*pb.LoginResponse, error) {
	request := &pb.LoginRequest{
		Login:    user.Login,
		Password: user.Password,
	}
	response, err := g.grpcClient.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GRPCClient) AddTextData(ctx context.Context, textData models.TextData, userSecretPhrase string) error {
	textData, err := g.hashService.HashTextData(textData, userSecretPhrase)
	if err != nil {
		return err
	}
	request := &pb.TextDataRequest{
		TextData:   textData.TextData,
		UserID:     textData.UserID,
		Metainfo:   textData.Metainfo,
		UploadedAt: timestamppb.New(time.Now()),
		UniqueKey:  random.GenerateUniqueKey(),
	}
	_, err = g.grpcClient.AddTextData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddTextData(ctx, textData)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCClient) AddBinaryData(ctx context.Context, binaryData models.BinaryData, userSecretPhrase string) error {
	binaryData, err := g.hashService.HashBinaryData(binaryData, userSecretPhrase)
	if err != nil {
		return err
	}
	request := &pb.BinaryDataRequest{
		BinaryData: binaryData.BinaryData,
		UserID:     binaryData.UserID,
		Metainfo:   binaryData.Metainfo,
		UploadedAt: timestamppb.New(time.Now()),
		UniqueKey:  random.GenerateUniqueKey(),
	}
	_, err = g.grpcClient.AddBinaryData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddBinaryData(ctx, binaryData)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCClient) AddCardData(ctx context.Context, cardData models.CardData, userSecretPhrase string) error {
	cardData, err := g.hashService.HashCardData(cardData, userSecretPhrase)
	if err != nil {
		return err
	}
	request := &pb.CardDataRequest{
		CardNumberData: cardData.CardNumberData,
		CardNameData:   cardData.CardNameData,
		CardDateData:   cardData.CardDateData,
		CvvData:        cardData.CvvData,
		UserID:         cardData.UserID,
		Metainfo:       cardData.Metainfo,
		UploadedAt:     timestamppb.New(time.Now()),
		UniqueKey:      random.GenerateUniqueKey(),
	}
	_, err = g.grpcClient.AddCardData(ctx, request)
	if err != nil {
		return err
	}
	err = g.saveService.AddCardData(ctx, cardData)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCClient) SyncData(ctx context.Context, userID int) error {
	uniqueKeys, err := g.getService.GetUniqueKeys(ctx, userID)
	if err != nil {
		g.log.Info(err)
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
	return nil
}
