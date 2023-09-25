package server_files_s3_service

import (
	"io"
	"os"
	server_files_s3_service "simple-hosting/controller/app/services/server-files-service/s3"
	"simple-hosting/controller/app/settings"
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"
	"testing"
)

func setEnvironment(config settings.ServiceSettings) {
	for k, v := range config.EnvironmentVars {
		os.Setenv(k, v)
	}
}

func TestSendingFileToS3(t *testing.T) {
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	setEnvironment(config)
	resp, err := server_files_s3_service.PublishFileToS3(server_files_s3_service.PublishFileToS3Request{
		Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
		Bucket:    "admin",
		ObjectKey: "admin/local_tests_file",
		FilePath:  "payload_transmit",
	})
	if err != nil {
		t.Error(err)
	}
	if !resp.Success {
		t.Error("Response must be success")
	}
}

func TestPullingFileFromS3(t *testing.T) {
	expectedText := "Hello, S3!"
	config := file_settings_provider.GetSetting[settings.ServiceSettings]("../../settings.yml")
	setEnvironment(config)
	_, err := server_files_s3_service.SaveFileFromS3(server_files_s3_service.SaveFileFromS3Request{
		Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
		Bucket:    "admin",
		ObjectKey: "admin/local_tests_file",
		FilePath:  "payload_save",
	})
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open("payload_save")
	if err != nil {
		t.Error(err)
	}
	buff, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	actual := string(buff)
	if actual != expectedText {
		t.Errorf("payload = %v, expected = %v", actual, expectedText)
	}
}
