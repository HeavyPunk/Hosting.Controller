package server_logging_service

type GetLogsOnPageRequest struct {
	Page int
}

type GetLogsOnPageResponse struct {
	Logs []struct {
		Id     int
		Record string
	}
}

type GetLatestLogsRequest struct {
	Page int
}

type GetLatestLogsResponse struct {
	Logs []struct {
		Id     int
		Record string
	}
}
