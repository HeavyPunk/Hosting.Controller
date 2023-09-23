package server_logging_file_service

type GetLogsOnPageRequest struct {
	Page int
}

type GetLogsOnPageResponse struct {
	Logs []struct {
		Id     int
		Record string
	}
}
