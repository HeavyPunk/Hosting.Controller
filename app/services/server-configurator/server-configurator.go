package server_configurator

import (
	file_settings_provider "simple-hosting/go-commons/settings/file-settings-provider"
)

func GetServerConfiguration(configPath string) ServerConfiguration {
	serverConfig := file_settings_provider.GetSetting[ServerConfiguration](configPath)
	return serverConfig
}
