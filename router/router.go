package router

import (
	"wanner-serve/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	gin.SetMode(viper.GetString("App_mode"))
	r := gin.Default()
	r.Use(middleware.Cros())
	router := r.Group("")
	{
		router.GET("/")
	}

	r.Run(viper.GetString("server.Http_port"))
}
