package main

import (
	"fmt"
	"os"
	files_controller "simple-hosting/controller/app/controllers/files-controller"
	server_controller "simple-hosting/controller/app/controllers/server-controller"
	state_controller "simple-hosting/controller/app/controllers/state"
	server_files_service "simple-hosting/controller/app/services/server-files-service"
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

	filesGroup := def.Group("/files")
	{
		filesGroup.POST("/s3/save-file", files_controller.GetFileS3)
		filesGroup.POST("/s3/push-file", files_controller.PushFileS3)
		filesGroup.POST("/delete-file", files_controller.RemoveFile)
		filesGroup.POST("/create-file", files_controller.CreateFile)
		filesGroup.POST("/create-directory", files_controller.CreateDirectory)
		filesGroup.POST("/list-directory", files_controller.ListDirectory)
		filesGroup.POST("/get-file-content-base64", files_controller.GetFileContentBase64)
		filesGroup.POST("/task-status", files_controller.PollTask)
		filesGroup.POST("/accept-task", files_controller.AcceptTask)
	}

	serviceGroup := def.Group("/state")
	{
		serviceGroup.GET("/_ping", state_controller.Ping)
		serverGroup.GET("/is-server-running")
	}

	server_files_service.Init(config)

	def.Run(":" + fmt.Sprint(config.App.Port))
}

func setEnvironment(config settings.ServiceSettings) {
	for k, v := range config.EnvironmentVars {
		os.Setenv(k, v)
	}
}
