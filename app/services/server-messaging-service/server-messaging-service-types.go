package server_messaging_service

type PostMessageRequest struct {
	PostSystem string `json:"post-system`
	Message    string `json:"message"`
}

type PostMessageResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"`
	Success  bool   `json:"success"`
}
