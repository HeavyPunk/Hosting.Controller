package server_controller

type StartServerRequest struct {
	SaveStdout bool `json:"save-stdout"`
	SaveStderr bool `json:"save-stderr"`
}

type StartServerResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
