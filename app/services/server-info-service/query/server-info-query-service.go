package server_info_query_service

import (
	"fmt"
	"simple-hosting/controller/app/settings"

	error_utils "simple-hosting/go-commons/tools/errors"

	"github.com/xrjr/mcutils/pkg/query"
)

func GetServerInfo() (ServerInfoQueryResponse, error) {
	config, err := settings.GetSettings()
	if err != nil {
		fmt.Printf("Error getting settings: %v\n", err)
		return ServerInfoQueryResponse{}, err
	}

	stat, err := query.QueryFull("localhost", int(config.App.Services.ServerInfo.Query.ServerQueryPort))
	if err != nil {
		fmt.Printf("Error getting server info: %v\n", err)
		return ServerInfoQueryResponse{}, err
	}

	onlinePlayers := stat.OnlinePlayers
	if onlinePlayers == nil {
		onlinePlayers = make([]string, 0)
	}

	return ServerInfoQueryResponse{
		OnlinePlayers: onlinePlayers,
		Properties:    stat.Properties,
		Error:         error_utils.GetErrorStringOrDefault(err, ""),
		Success:       err == nil,
	}, nil
}
