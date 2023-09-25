package server_files_service_test

import (
	"os"
	server_files_service "simple-hosting/controller/app/services/server-files-service"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"
	"testing"
	"time"
)

func setEnvironment(config settings.ServiceSettings) {
	for k, v := range config.EnvironmentVars {
		os.Setenv(k, v)
	}
}

func TestTransmitFileViaS3(t *testing.T) {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	setEnvironment(config)
	server_files_service.Init(config)
	taskId, err := server_files_service.TransferFile(server_files_service.TransferFileRequest{
		System: "S3",
		S3: struct {
			Endpoint  string
			Bucket    string
			ObjectKey string
			FilePath  string
		}{
			Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
			Bucket:    "admin",
			ObjectKey: "admin/local_tests_files_transmit",
			FilePath:  "payload_transmit",
		},
	})
	if err != nil {
		t.Error(err)
	}

	attempt := 0
	for {
		if attempt == 5 {
			t.Errorf("Timeout: %v attempt", attempt)
			break
		}
		time.Sleep(1 * time.Second)
		resp, err := server_files_service.PollTask(taskId.TaskId)
		if err != nil {
			t.Error(err)
		}
		if resp.TaskStatus == server_files_service.Failed {
			t.Errorf("Expected completed Task, got %v. Error: %v", resp.TaskStatus, resp.Error)
			break
		}
		if resp.TaskStatus == server_files_service.Completed {
			break
		}
		attempt++
	}
}

func TestSaveFileFromS3(t *testing.T) {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	setEnvironment(config)
	server_files_service.Init(config)
	taskId, err := server_files_service.SaveFile(server_files_service.SaveFileRequest{
		System: "S3",
		S3: struct {
			Endpoint  string
			Bucket    string
			ObjectKey string
			FilePath  string
		}{
			Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
			Bucket:    "admin",
			ObjectKey: "admin/local_tests_files_transmit",
			FilePath:  "payload_save",
		},
	})
	if err != nil {
		t.Error(err)
	}

	attempt := 0
	for {
		if attempt == 5 {
			t.Errorf("Timeout: %v attempt", attempt)
			break
		}
		time.Sleep(1 * time.Second)
		resp, err := server_files_service.PollTask(taskId.TaskId)
		if err != nil {
			t.Error(err)
		}
		if resp.TaskStatus == server_files_service.Failed {
			t.Errorf("Expected completed Task, got %v. Error: %v", resp.TaskStatus, resp.Error)
			break
		}
		if resp.TaskStatus == server_files_service.Completed {
			break
		}
		attempt++
	}
}
