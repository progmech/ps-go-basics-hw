package main

import (
	"json-loader/bins"
	"json-loader/storage"
)

func main() {
	var binList bins.BinList = bins.BinList{}
	bin := bins.NewBin("a", true, "Test1")
	binList = append(binList, *bin)
	storage.SaveBins(&binList)
	storage.PrintBins(&binList)
}
