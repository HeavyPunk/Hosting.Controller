package main

import (
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
		gr.POST("/start")
	}
}
