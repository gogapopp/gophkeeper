package service

import (
	"context"
	"fmt"
	"os"

	"github.com/gogapopp/gophkeeper/internal/hasher"
	"github.com/gogapopp/gophkeeper/models"
)

// интерфейс взаимодействия с БД для получения данных
type Getter interface {
	GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error)
	GetTextData(ctx context.Context, uniqueKey int) (models.TextData, error)
	GetBinaryData(ctx context.Context, uniqueKey int) (models.BinaryData, error)
	GetCardData(ctx context.Context, uniqueKey int) (models.CardData, error)
	GetDatas(ctx context.Context, table string) (map[int]string, error)
}

// GetUniqueKeys получает все уникальные ключи ключи пользователя для каждого типа данных
func (g *GetService) GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error) {
	const op = "service.getter.GetUniqueKeys"
	uniqueKeys, err := g.get.GetUniqueKeys(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return uniqueKeys, nil
}

// GetTextData получает текстовые данные пользователя по уникальному ключу и дешефрует данные
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
	encryptedTextData, err := hasher.Decrypt(textdata.TextData, []byte(userSecretPhrase))
	if err != nil {
		return textdata, err
	}
	encryptedMetainfo, err := hasher.Decrypt(textdata.Metainfo, []byte(userSecretPhrase))
	if err != nil {
		return textdata, err
	}
	// пишем в файл
	_, err = fmt.Fprintf(file, "UniqueKey: %s\nTextData: %s\nUploadedAt: %s\nMetainfo: %s\n",
		textdata.UniqueKey, string(encryptedTextData), uploadedAt, string(encryptedMetainfo))
	if err != nil {
		return textdata, fmt.Errorf("%s: %s", op, err)
	}
	return textdata, nil
}

// GetBinaryData получает бинарные данные пользователя по уникальному ключу и дешефрует данные
func (g *GetService) GetBinaryData(ctx context.Context, uniqueKey int, userSecretPhrase string) (models.BinaryData, error) {
	const op = "service.getter.GetBinaryData"
	binarydata, err := g.get.GetBinaryData(ctx, uniqueKey)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	fileTxtName := fmt.Sprintf("binary_data_%d.txt", uniqueKey)
	fileTXT, err := os.Create(fileTxtName)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	defer fileTXT.Close()
	fileBinaryName := fmt.Sprintf("binary_data_%d.exe", uniqueKey)
	fileEXE, err := os.Create(fileBinaryName)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	defer fileEXE.Close()
	encryptedBinary, err := hasher.Decrypt(binarydata.BinaryData, []byte(userSecretPhrase))
	if err != nil {
		return binarydata, err
	}
	// создаём файл .exe
	err = os.WriteFile(fileBinaryName, encryptedBinary, 0644)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	uploadedAt := binarydata.UploadedAt.AsTime()
	encryptedMetainfo, err := hasher.Decrypt(binarydata.Metainfo, []byte(userSecretPhrase))
	if err != nil {
		return binarydata, err
	}
	// пишем в файл .txt
	_, err = fmt.Fprintf(fileTXT, "UniqueKey: %d\nUploadedAt: %s\nMetainfo: %s\n",
		uniqueKey, uploadedAt, encryptedMetainfo)
	if err != nil {
		return binarydata, fmt.Errorf("%s: %s", op, err)
	}
	return binarydata, nil
}

// GetCardData получает данные карты пользователя по уникальному ключу и дешефрует данные
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
	// дешефруем данные
	encryptedCardNumber, err := hasher.Decrypt(carddata.CardNumberData, []byte(userSecretPhrase))
	if err != nil {
		return carddata, err
	}
	encryptedName, err := hasher.Decrypt(carddata.CardNameData, []byte(userSecretPhrase))
	if err != nil {
		return carddata, err
	}
	encryptedCardDate, err := hasher.Decrypt(carddata.CardDateData, []byte(userSecretPhrase))
	if err != nil {
		return carddata, err
	}
	encryptedCVV, err := hasher.Decrypt(carddata.CvvData, []byte(userSecretPhrase))
	if err != nil {
		return carddata, err
	}
	encryptedMetainfo, err := hasher.Decrypt(carddata.Metainfo, []byte(userSecretPhrase))
	if err != nil {
		return carddata, err
	}
	// пишем в файл
	_, err = fmt.Fprintf(file, "UniqueKey: %s\nCardNumberData: %s\nCardNameData: %s\nCardDateData: %s\nCVVData: %s\nUploadedAt: %s\nMetainfo: %s\n",
		carddata.UniqueKey, encryptedCardNumber, encryptedName, encryptedCardDate, encryptedCVV, uploadedAt, encryptedMetainfo)
	if err != nil {
		return carddata, fmt.Errorf("%s: %s", op, err)
	}
	return carddata, nil
}

// GetDatas получает список уникальных ключей и даты сохранения каждой строки определёного типа данных
func (g *GetService) GetDatas(ctx context.Context, table string) (map[int]string, error) {
	const op = "service.getter.GetDatas"
	datas, err := g.get.GetDatas(ctx, table)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return datas, nil
}
