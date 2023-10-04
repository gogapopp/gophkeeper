package usecase

import (
	"context"

	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
)

type Getter interface {
	GetDatas(uniqueKeys map[string][]string) (models.SyncData, error)
}

func (g *GetUsecase) GetDatas(ctx context.Context, in *pb.SyncRequest) (*pb.SyncResponse, error) {
	uniqueKeys := make(map[string][]string)
	for k, v := range in.Keys {
		uniqueKeys[k] = v.Values
	}
	datas, err := g.get.GetDatas(uniqueKeys)
	if err != nil {
		return nil, err
	}
	var textData []*pb.TextDataRequest
	for _, td := range datas.TextData {
		textData = append(textData, &pb.TextDataRequest{
			TextData:   td.TextData,
			Metainfo:   td.Metainfo,
			UserID:     td.UserID,
			UniqueKey:  td.UniqueKey,
			UploadedAt: td.UploadedAt,
		})
	}
	var binaryData []*pb.BinaryDataRequest
	for _, bd := range datas.BinaryData {
		binaryData = append(binaryData, &pb.BinaryDataRequest{
			BinaryData: bd.BinaryData,
			Metainfo:   bd.Metainfo,
			UserID:     bd.UserID,
			UniqueKey:  bd.UniqueKey,
			UploadedAt: bd.UploadedAt,
		})
	}
	var cardData []*pb.CardDataRequest
	for _, cd := range datas.CardData {
		cardData = append(cardData, &pb.CardDataRequest{
			CardNumberData: cd.CardNumberData,
			CardNameData:   cd.CardNameData,
			CardDateData:   cd.CardDateData,
			CvvData:        cd.CvvData,
			Metainfo:       cd.Metainfo,
			UserID:         cd.UserID,
			UniqueKey:      cd.UniqueKey,
			UploadedAt:     cd.UploadedAt,
		})
	}
	responseData := &pb.SyncResponse{
		TextData:   textData,
		BinaryData: binaryData,
		CardData:   cardData,
	}
	return responseData, nil
}
