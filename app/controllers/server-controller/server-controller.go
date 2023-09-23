package server_controller

import (
	"fmt"
	"net/http"
	server_configurator "simple-hosting/controller/app/services/server-configurator"
	server_info_service "simple-hosting/controller/app/services/server-info-service"
	server_logging_service "simple-hosting/controller/app/services/server-logging-service"
	server_manager "simple-hosting/controller/app/services/server-manager"
	server_messaging_service "simple-hosting/controller/app/services/server-messaging-service"

	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"
	errors_tools "simple-hosting/go-commons/tools/errors"

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
	}, config)

	c.JSON(http.StatusOK, StartServerResponse{
		Success: resp.Success,
		Error:   errors_tools.GetErrorStringOrDefault(resp.Error, ""),
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
		Error:   errors_tools.GetErrorStringOrDefault(resp.Error, ""),
	})
}

func IsServerRunning(c *gin.Context) {
	serverManager := server_manager.Init()
	res, err := serverManager.CheckForServerRunning()
	c.JSON(http.StatusOK, IsServerRunningResponse{
		IsRunning: res,
		Success:   err == nil,
		Error:     errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func PostMessageToServer(c *gin.Context) {
	var request PostMessageToServerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp := server_messaging_service.PostMessage(server_messaging_service.PostMessageRequest{
		PostSystem: request.PostSystem,
		Message:    request.Message,
	})

	c.JSON(http.StatusOK, PostMessageToServerResponse{
		Response: resp.Response,
		Success:  resp.Success,
		Error:    resp.Error,
	})
}

func GetServerInfo(c *gin.Context) {
	var request GetServerInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_info_service.GetServerInfo(server_info_service.GetServerInfoRequest{
		PostSystem: request.PostSystem,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GetServerInfo failed with error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetServerInfoResponse{
		OnlinePlayers: resp.OnlinePlayers,
		Properties:    resp.Properties,
		Error:         resp.Error,
		Success:       resp.Success,
	})
}

func GetServerLogsOnPage(c *gin.Context) {
	config, err := settings.GetSettings()
	if err != nil {
		fmt.Printf("Error getting settings: %v", err)
		c.JSON(http.StatusInternalServerError, GetServerLogsOnPageResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	var request GetServerLogsOnPageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, GetServerLogsOnPageResponse{
			Error:   err.Error(),
			Success: false,
		})
		return
	}

	resp, err := server_logging_service.GetLogsOnPage(
		server_logging_service.GetLogsOnPageRequest{
			Page: request.Page,
		},
		config,
	)

	c.JSON(http.StatusOK, GetServerLogsOnPageResponse{
		Logs:    resp.Logs,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
		Success: err == nil,
	})
}

func GetServerLogsOnLastPage(c *gin.Context) {
	config, err := settings.GetSettings()
	if err != nil {
		fmt.Printf("Error getting settings: %v", err)
		c.JSON(http.StatusInternalServerError, GetServerLogsOnPageResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	resp, err := server_logging_service.GetLatestLogs(config)

	c.JSON(http.StatusOK, GetServerLogsOnPageResponse{
		Logs:    resp.Logs,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
		Success: err == nil,
	})
}
