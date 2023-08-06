package main

import (
	"wanner-serve/config"
	"wanner-serve/model"
	"wanner-serve/router"
)

func main() {
	config.InitConfig()
	model.InitDb()
	router.InitRouter()
}
