package main

import (
	"fmt"
	"os"
	files_controller "simple-hosting/controller/app/controllers/files-controller"
	server_controller "simple-hosting/controller/app/controllers/server-controller"
	state_controller "simple-hosting/controller/app/controllers/state"
	"simple-hosting/controller/app/settings"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := settings.GetSettings()
	if err != nil {
		fmt.Printf("Error getting settings: %v\n", err)
		return
	}
	setEnvironment(config)
	gin.SetMode(config.App.Configuration)
	def := gin.Default()
	serverGroup := def.Group("/server")
	{
		serverGroup.POST("/start", server_controller.StartServer)
		serverGroup.POST("/stop", server_controller.StopServer)
		serverGroup.POST("/messaging/post", server_controller.PostMessageToServer)
		serverGroup.POST("/info", server_controller.GetServerInfo)
		serverGroup.GET("/is-running", server_controller.IsServerRunning)
		serverGroup.POST("/logs/get-server-log-on-page", server_controller.GetServerLogsOnPage)
		serverGroup.POST("/logs/get-server-last-log", server_controller.GetServerLogsOnLastPage)
	}

	filesGroup := def.Group("/files/s3")
	{
		filesGroup.POST("/save-file", files_controller.GetFileS3)
		filesGroup.POST("/push-file", files_controller.PushFileS3)
		filesGroup.POST("/task-status", files_controller.PollTask)
	}

	serviceGroup := def.Group("/state")
	{
		serviceGroup.GET("/_ping", state_controller.Ping)
		serverGroup.GET("/is-server-running")
	}

	def.Run(":" + fmt.Sprint(config.App.Port))
}

func setEnvironment(config settings.ServiceSettings) {
	for k, v := range config.EnvironmentVars {
		os.Setenv(k, v)
	}
}
