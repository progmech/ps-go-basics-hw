package main

import (
	"json-loader/api"
	"json-loader/bins"
	"json-loader/storage"
)

func main() {
	binList := bins.BinList{
		*bins.NewBin("a", true, "Test1"),
		*bins.NewBin("b", true, "Test2"),
		*bins.NewBin("c", true, "Test3"),
	}
	var store storage.Storager = storage.NewFileStorage("storage.json", &binList)
	store.Write()
	store.Read()

	api.ReadKey()
}
