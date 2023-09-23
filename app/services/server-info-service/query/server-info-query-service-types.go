package server_info_query_service

type ServerInfoQueryResponse struct {
	OnlinePlayers []string          `json:"online-players"`
	Properties    map[string]string `json:"properties"`
	Error         string            `json:"error"`
	Success       bool              `json:"success"`
}
