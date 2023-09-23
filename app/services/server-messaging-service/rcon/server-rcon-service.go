package server_rcon_service

import (
	"fmt"
	"simple-hosting/controller/app/settings"

	"github.com/xrjr/mcutils/pkg/rcon"
)

func Execute(input ExecuteInput) (ExecuteResponse, error) {
	config, err := settings.GetSettings()
	if err != nil {
		fmt.Printf("Error getting settings: %v\n", err)
		return ExecuteResponse{}, err
	}

	resp, err := rcon.Rcon(
		"localhost",
		int(config.App.Services.ServerMessaging.Rcon.ServerRconPort),
		config.App.Services.ServerMessaging.Rcon.ServerRconPassword,
		input.Command,
	)

	if err != nil {
		fmt.Printf("Error creating rcon connection: %v\n", err)
		return ExecuteResponse{}, err
	}

	return ExecuteResponse{Response: resp}, nil
}
