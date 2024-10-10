package main

import (
	"json-loader/bins"
	"json-loader/storage"
)

func main() {
	binList := bins.BinList{
		*bins.NewBin("a", true, "Test1"),
		*bins.NewBin("b", true, "Test2"),
		*bins.NewBin("c", true, "Test3"),
	}
	storage := storage.NewFileStorage("storage.json", &binList)
	storage.Write()
	storage.Read()
}
