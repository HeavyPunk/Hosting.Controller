package server_logging_file_service_tests

import (
	server_logging_file_service "simple-hosting/controller/app/services/server-logging-service/server-logging-file-service"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"
	"testing"
)

func TestReadFirstPage(t *testing.T) {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	logs, err := server_logging_file_service.GetLogsOnPage(server_logging_file_service.GetLogsOnPageRequest{
		Page: 0,
	}, config)
	if err != nil {
		t.Error(err)
		return
	}
	if len(logs.Logs) == 0 {
		t.Errorf("Page size is not equal to server logging page size (%d got %d)", config.App.Services.ServerLogging.PageSize, len(logs.Logs))
	}
}

func TestReadLastPage(t *testing.T) {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	logs, err := server_logging_file_service.GetLogsOnLastPage(config)
	if err != nil {
		t.Error(err)
		return
	}
	if len(logs.Logs) == 0 {
		t.Errorf("Page size is not equal to server logging page size (%d got %d)", config.App.Services.ServerLogging.PageSize, len(logs.Logs))
	}
}
