package storage

import (
	"encoding/json"
	"fmt"
	"json-loader/bins"
	"json-loader/file"
)

type Storager interface {
	Read()
	Write()
}

type FileStorage struct {
	Bins     bins.BinList `json:"bins"`
	filePath string
}

func NewFileStorage(path string, binList *bins.BinList) *FileStorage {
	return &FileStorage{
		Bins:     *binList,
		filePath: path,
	}
}

func (fs *FileStorage) Read() {
	data, err := file.ReadFile(fs.filePath)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %s.\n", fs.filePath, err)
		return
	}

	err = json.Unmarshal(data, &fs.Bins)
	if err != nil {
		fmt.Printf("Ошибка преобразования файла %s: %s.\n", fs.filePath, err)
		return
	}

	for _, v := range fs.Bins {
		fmt.Println(v)
	}
}

func (fs *FileStorage) Write() {
	data, err := json.Marshal(fs.Bins)
	if err != nil {
		fmt.Printf("Ошибка преобразования файла %s: %s.\n", fs.filePath, err)
		return
	}

	err = file.SaveFile(fs.filePath, data)
	if err != nil {
		fmt.Printf("Ошибка сохранения файла %s: %s.\n", fs.filePath, err)
	}

	fmt.Println("Файл успешно сохранён.")
}
