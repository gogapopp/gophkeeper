package service

import (
	"context"
	"fmt"
	"os"

	"github.com/gogapopp/gophkeeper/models"
)

type Getter interface {
	GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error)
	GetTextData(ctx context.Context, uniqueKey int) (models.TextData, error)
	GetBinaryData(ctx context.Context, uniqueKey int) (models.BinaryData, error)
	GetCardData(ctx context.Context, uniqueKey int) (models.CardData, error)
	GetDatas(ctx context.Context, table string) (map[int]string, error)
}

func (g *GetService) GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error) {
	const op = "service.getter.GetUniqueKeys"
	uniqueKeys, err := g.get.GetUniqueKeys(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return uniqueKeys, nil
}

func (g *GetService) GetTextData(ctx context.Context, uniqueKey int, userSecretPhrase string) (models.TextData, error) {
	const op = "service.getter.GetTextData"
	textdata, err := g.get.GetTextData(ctx, uniqueKey)
	if err != nil {
		return textdata, fmt.Errorf("%s: %s", op, err)
	}
	fileName := fmt.Sprintf("text_data_%d.txt", uniqueKey)
	file, err := os.Create(fileName)
	if err != nil {
		return textdata, fmt.Errorf("%s: %s", op, err)
	}
	defer file.Close()
	uploadedAt := textdata.UploadedAt.AsTime()
	// TODO: расшифровку добавить
	// encryptedTextData, err := hash.Decrypt(textdata.TextData, []byte(userSecretPhrase))
	// if err != nil {
	// 	return textdata, err
	// }
	_, err = fmt.Fprintf(file, "UniqueKey: %s\nTextData: %s\nUploadedAt: %s\nMetainfo: %s\n",
		textdata.UniqueKey, textdata.TextData, uploadedAt, textdata.Metainfo)
	if err != nil {
		return textdata, fmt.Errorf("%s: %s", op, err)
	}
	return textdata, nil
}

func (g *GetService) GetBinaryData(ctx context.Context, uniqueKey int, userSecretPhrase string) (models.BinaryData, error) {
	const op = "service.getter.GetBinaryData"
	binarydata, err := g.get.GetBinaryData(ctx, uniqueKey)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	fileName := fmt.Sprintf("binary_data_%d", uniqueKey)
	file, err := os.Create(fileName)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	defer file.Close()
	err = os.WriteFile(fileName, binarydata.BinaryData, 0644)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	uploadedAt := binarydata.UploadedAt.AsTime()
	// TODO: расшифровку добавить
	_, err = fmt.Fprintf(file, "UniqueKey: %s\nUploadedAt: %s\nMetainfo: %s\n",
		binarydata.UniqueKey, uploadedAt, binarydata.Metainfo)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	return binarydata, nil
}

func (g *GetService) GetCardData(ctx context.Context, uniqueKey int, userSecretPhrase string) (models.CardData, error) {
	const op = "service.getter.GetCardData"
	carddata, err := g.get.GetCardData(ctx, uniqueKey)
	if err != nil {
		return carddata, fmt.Errorf("%s: %s", op, err)
	}
	fileName := fmt.Sprintf("card_data_%d.txt", uniqueKey)
	file, err := os.Create(fileName)
	if err != nil {
		return carddata, fmt.Errorf("%s: %s", op, err)
	}
	defer file.Close()
	uploadedAt := carddata.UploadedAt.AsTime()
	// TODO: расшифровку добавить
	_, err = fmt.Fprintf(file, "UniqueKey: %s\nCardNumberData: %s\nCardNameData: %s\nCardDateData: %s\nCVVData: %s\nUploadedAt: %s\nMetainfo: %s\n",
		carddata.UniqueKey, carddata.CardNumberData, carddata.CardNameData, carddata.CardDateData, carddata.CvvData, uploadedAt, carddata.Metainfo)
	if err != nil {
		return carddata, fmt.Errorf("%s: %s", op, err)
	}
	return carddata, nil
}

func (g *GetService) GetDatas(ctx context.Context, table string) (map[int]string, error) {
	const op = "service.getter.GetDatas"
	datas, err := g.get.GetDatas(ctx, table)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return datas, nil
}
