package server_info_service

import (
	"fmt"
	server_info_query_service "simple-hosting/controller/app/services/server-info-service/query"
	error_utils "simple-hosting/go-commons/tools/errors"
)

func GetServerInfo(request GetServerInfoRequest) (GetServerInfoResponse, error) {
	switch request.PostSystem {
	case "query":
		resp, err := server_info_query_service.GetServerInfo()
		if err != nil {
			fmt.Print("Error getting server info: %v\n", err)
			return GetServerInfoResponse{}, err
		}
		return GetServerInfoResponse{
			OnlinePlayers: resp.OnlinePlayers,
			Properties:    resp.Properties,
			Error:         error_utils.GetErrorStringOrDefault(err, resp.Error),
			Success:       resp.Success,
		}, nil
	default:
		return GetServerInfoResponse{}, fmt.Errorf("Unknown post system: %v", request.PostSystem)
	}
}
