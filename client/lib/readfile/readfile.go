package readfile

import "os"

func ReadFile(filePath string) ([]byte, error) {
	// Чтение файла в массив байтов
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
