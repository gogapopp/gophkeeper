package service

import (
	"github.com/gogapopp/gophkeeper/internal/hasher"
	"github.com/gogapopp/gophkeeper/models"
)

func (h *HashService) HashTextData(textdata models.TextData, userSecretPhrase string) (models.TextData, error) {
	hashText, err := hasher.Encrypt([]byte(textdata.TextData), []byte(userSecretPhrase))
	if err != nil {
		return models.TextData{}, err
	}
	hashMetainfo, err := hasher.Encrypt([]byte(textdata.Metainfo), []byte(userSecretPhrase))
	if err != nil {
		return models.TextData{}, err
	}
	hashTextData := models.TextData{
		TextData: hashText,
		Metainfo: hashMetainfo,
		UserID:   textdata.UserID,
	}
	return hashTextData, nil
}

func (h *HashService) HashBinaryData(binarydata models.BinaryData, userSecretPhrase string) (models.BinaryData, error) {
	hashBinary, err := hasher.Encrypt([]byte(binarydata.BinaryData), []byte(userSecretPhrase))
	if err != nil {
		return models.BinaryData{}, err
	}
	hashMetainfo, err := hasher.Encrypt([]byte(binarydata.Metainfo), []byte(userSecretPhrase))
	if err != nil {
		return models.BinaryData{}, err
	}
	hashBinaryData := models.BinaryData{
		BinaryData: hashBinary,
		Metainfo:   hashMetainfo,
		UserID:     binarydata.UserID,
	}
	return hashBinaryData, nil
}
func (h *HashService) HashCardData(carddata models.CardData, userSecretPhrase string) (models.CardData, error) {
	hashCardNumber, err := hasher.Encrypt([]byte(carddata.CardNumberData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardName, err := hasher.Encrypt([]byte(carddata.CardNameData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardDate, err := hasher.Encrypt([]byte(carddata.CardDateData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardCVV, err := hasher.Encrypt([]byte(carddata.CvvData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashMetainfo, err := hasher.Encrypt([]byte(carddata.Metainfo), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardData := models.CardData{
		CardNumberData: hashCardNumber,
		CardNameData:   hashCardName,
		CardDateData:   hashCardDate,
		CvvData:        hashCardCVV,
		Metainfo:       hashMetainfo,
		UserID:         carddata.UserID,
	}
	return hashCardData, nil
}
