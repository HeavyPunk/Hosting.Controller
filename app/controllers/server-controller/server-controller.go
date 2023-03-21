package server_controller

import (
	"net/http"
	server_configurator "simple-hosting/controller/app/services/server-configurator"
	server_manager "simple-hosting/controller/app/services/server-manager"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"

	"github.com/gin-gonic/gin"
)

func StartServer(c *gin.Context) {
	var request StartServerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("settings.yml")
	serverManager := server_manager.Init()
	serverConfig := server_configurator.GetServerConfiguration(config.Server.StartupConfig)
	resp := serverManager.RunServer(server_manager.RunServerRequest{
		RunCmd:          serverConfig.RunCmd.Cmd,
		Args:            serverConfig.RunCmd.Args,
		WorkingDir:      serverConfig.WorkingDir,
		EnvironmentVars: serverConfig.EnvVars,
		SaveStdout:      request.SaveStdout,
		SaveStderr:      request.SaveStderr,
	})
	var errStr string
	if resp.Error != nil {
		errStr = resp.Error.Error()
	}
	c.JSON(http.StatusOK, StartServerResponse{
		Success: resp.Success,
		Error:   errStr,
	})
}

func StopServer(c *gin.Context) {
	var request StopServerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}
	serverManager := server_manager.Init()
	resp := serverManager.StopServer(server_manager.StopServerRequest{
		ForceInterrupt: request.Force,
	})
	c.JSON(http.StatusOK, StopServerResponse{
		Success: resp.Success,
		Error:   resp.Error.Error(),
	})
}
