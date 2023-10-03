package service

import (
	"github.com/gogapopp/gophkeeper/internal/hash"
	"github.com/gogapopp/gophkeeper/models"
)

func (h *HashService) HashTextData(textdata models.TextData, userSecretPhrase string) (models.TextData, error) {
	hashText, err := hash.Encrypt([]byte(textdata.TextData), []byte(userSecretPhrase))
	if err != nil {
		return models.TextData{}, err
	}
	hashMetainfo, err := hash.Encrypt([]byte(textdata.Metainfo), []byte(userSecretPhrase))
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
	hashBinary, err := hash.Encrypt([]byte(binarydata.BinaryData), []byte(userSecretPhrase))
	if err != nil {
		return models.BinaryData{}, err
	}
	hashMetainfo, err := hash.Encrypt([]byte(binarydata.Metainfo), []byte(userSecretPhrase))
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
	hashCardNumber, err := hash.Encrypt([]byte(carddata.CardNumberData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardName, err := hash.Encrypt([]byte(carddata.CardNameData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardDate, err := hash.Encrypt([]byte(carddata.CardDateData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashCardCVV, err := hash.Encrypt([]byte(carddata.CvvData), []byte(userSecretPhrase))
	if err != nil {
		return models.CardData{}, err
	}
	hashMetainfo, err := hash.Encrypt([]byte(carddata.Metainfo), []byte(userSecretPhrase))
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
