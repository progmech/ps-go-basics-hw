package api

import (
	"fmt"
	"json-loader/config"
)

func ReadKey() {
	conf := config.NewConfig()
	fmt.Println(conf.Key)
}
