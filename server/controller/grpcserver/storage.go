package grpcserver

import (
	"context"

	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gs *grpcServer) AddTextData(ctx context.Context, in *pb.TextDataRequest) (*pb.Empty, error) {
	var response pb.Empty
	textdata := models.TextData{
		UserID:     in.UserID,
		UniqueKey:  in.UniqueKey,
		TextData:   in.TextData,
		UploadedAt: in.UploadedAt,
		Metainfo:   in.Metainfo,
	}
	err := gs.store.AddTextData(ctx, textdata)
	if err != nil {
		gs.log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	return &response, nil
}

func (gs *grpcServer) AddBinaryData(ctx context.Context, in *pb.BinaryDataRequest) (*pb.Empty, error) {
	var response pb.Empty
	binarydata := models.BinaryData{
		UserID:     in.UserID,
		UniqueKey:  in.UniqueKey,
		BinaryData: in.BinaryData,
		UploadedAt: in.UploadedAt,
		Metainfo:   in.Metainfo,
	}
	err := gs.store.AddBinaryData(ctx, binarydata)
	if err != nil {
		gs.log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	return &response, nil
}

func (gs *grpcServer) AddCardData(ctx context.Context, in *pb.CardDataRequest) (*pb.Empty, error) {
	var response pb.Empty
	carddata := models.CardData{
		UserID:         in.UserID,
		UniqueKey:      in.UniqueKey,
		CardNumberData: in.CardNumberData,
		CardNameData:   in.CardNameData,
		CardDateData:   in.CardDateData,
		CvvData:        in.CvvData,
		UploadedAt:     in.UploadedAt,
		Metainfo:       in.Metainfo,
	}
	err := gs.store.AddCardData(ctx, carddata)
	if err != nil {
		gs.log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	return &response, nil
}
