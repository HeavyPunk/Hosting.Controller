package server_messaging_service

import (
	"fmt"
	server_rcon_service "simple-hosting/controller/app/services/server-messaging-service/rcon"
	error_utils "simple-hosting/go-commons/tools/errors"
)

func PostMessage(request PostMessageRequest) PostMessageResponse {
	switch request.PostSystem {
	case "rcon":
		resp, err := server_rcon_service.Execute(server_rcon_service.ExecuteInput{
			Command: request.Message,
		})
		return PostMessageResponse{
			Response: resp.Response,
			Error:    error_utils.GetErrorStringOrDefault(err, ""),
			Success:  err == nil,
		}
	default:
		return PostMessageResponse{Success: false, Error: fmt.Sprintf("Unsupported post system %s", request.PostSystem)}
	}
}
