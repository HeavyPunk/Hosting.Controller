package main

import (
	"fmt"
	server_controller "simple-hosting/controller/app/controllers/server-controller"
	state_controller "simple-hosting/controller/app/controllers/state"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"

	"github.com/gin-gonic/gin"
)

func main() {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("settings.yml")
	gin.SetMode(config.App.Configuration)
	def := gin.Default()
	serverGroup := def.Group("/server")
	{
		serverGroup.POST("/start", server_controller.StartServer)
		serverGroup.POST("/stop", server_controller.StopServer)
	}

	serviceGroup := def.Group("/state")
	{
		serviceGroup.GET("/_ping", state_controller.Ping)
	}

	def.Run(":" + fmt.Sprint(config.App.Port))
}
