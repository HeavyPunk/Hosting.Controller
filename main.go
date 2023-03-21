package main

import (
	"fmt"
	server_controller "simple-hosting/controller/app/controllers/server-controller"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"

	"github.com/gin-gonic/gin"
)

func main() {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("settings.yml")
	gin.SetMode(config.App.Configuration)
	def := gin.Default()
	gr := def.Group("/server")
	{
		gr.POST("/start", server_controller.StartServer)
		gr.POST("/stop", server_controller.StopServer)
	}
	def.Run(":" + fmt.Sprint(config.App.Port))
}
