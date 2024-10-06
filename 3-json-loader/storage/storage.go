package storage

import (
	"encoding/json"
	"fmt"
	"json-loader/bins"
	"json-loader/file"
)

const storage = "storage.json"

func PrintBins(binList *bins.BinList) {
	data, err := file.ReadFile(storage)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %s.\n", storage, err)
		return
	}

	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Printf("Ошибка преобразования файла %s: %s.\n", storage, err)
		return
	}

	for _, v := range *binList {
		fmt.Println(v)
	}
}

func SaveBins(binList *bins.BinList) {
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Printf("Ошибка преобразования файла %s: %s.\n", storage, err)
		return
	}

	err = file.SaveFile(storage, data)
	if err != nil {
		fmt.Printf("Ошибка сохранения файла %s: %s.\n", storage, err)
	}

	fmt.Println("Файл успешно сохранён.")
}
