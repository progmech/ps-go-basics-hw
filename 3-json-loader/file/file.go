package file

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	if !isJsonFile(path) {
		return nil, errors.New("расширение файла не json")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func isJsonFile(path string) bool {
	fileParts := strings.Split(path, ".")
	return fileParts[len(fileParts)-1] == "json"
}

func SaveFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
