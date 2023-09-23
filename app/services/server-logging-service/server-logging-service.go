package server_logging_service

import (
	"fmt"
	server_logging_file_service "simple-hosting/controller/app/services/server-logging-service/server-logging-file-service"
	"simple-hosting/controller/app/settings"
)

func GetLogsOnPage(request GetLogsOnPageRequest, config settings.ServiceSettings) (GetLogsOnPageResponse, error) {
	switch config.App.Services.ServerLogging.LoggingSystem {
	case "file":
		res, err := server_logging_file_service.GetLogsOnPage(server_logging_file_service.GetLogsOnPageRequest{
			Page: request.Page,
		}, config)
		if err != nil {
			fmt.Printf("Error getting logs on page %d: %v", request.Page, err)
			return GetLogsOnPageResponse{}, err
		}
		return GetLogsOnPageResponse{Logs: res.Logs}, nil
	default:
		return GetLogsOnPageResponse{}, fmt.Errorf("unknown logging system requested %s", config.App.Services.ServerLogging.LoggingSystem)
	}
}

func GetLatestLogs(config settings.ServiceSettings) (GetLatestLogsResponse, error) {
	switch config.App.Services.ServerLogging.LoggingSystem {
	case "file":
		res, err := server_logging_file_service.GetLogsOnLastPage(config)
		if err != nil {
			fmt.Printf("Error getting logs on last page: %v", err)
			return GetLatestLogsResponse{}, err
		}
		return GetLatestLogsResponse{Logs: res.Logs}, nil
	default:
		return GetLatestLogsResponse{}, fmt.Errorf("unknown logging system requested %s", config.App.Services.ServerLogging.LoggingSystem)
	}
}
