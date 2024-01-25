package server_info_service

type GetServerInfoRequest struct {
	PostSystem string `json:"post-system"`
}

type GetServerInfoResponse struct {
	OnlinePlayers []string          `json:"onlinePlayers"`
	Properties    map[string]string `json:"properties"`
	Error         string            `json:"error"`
	Success       bool              `json:"success"`
}
