package server_controller

type StartServerRequest struct {
	SaveStdout bool `json:"save-stdout"`
	SaveStderr bool `json:"save-stderr"`
}

type StartServerResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type StopServerRequest struct {
	Force bool `json:"force"`
}

type StopServerResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type IsServerRunningResponse struct {
	IsRunning bool   `json:"is-running"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

type PostMessageToServerRequest struct {
	PostSystem string `json:"post-system"`
	Message    string `json:"message"`
}

type PostMessageToServerResponse struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
	Error    string `json:"error"`
}

type GetServerInfoRequest struct {
	PostSystem string `json:"post-system"`
}

type GetServerInfoResponse struct {
	OnlinePlayers []string          `json:"online-players"`
	Properties    map[string]string `json:"properties"`
	Error         string            `json:"error"`
	Success       bool              `json:"success"`
}

type GetServerLogsOnPageRequest struct {
	Page int `json:"page"`
}

type GetServerLogsOnPageResponse struct {
	Logs []struct {
		Id     int
		Record string
	}
	Error   string `json:"error"`
	Success bool   `json:"success"`
}
