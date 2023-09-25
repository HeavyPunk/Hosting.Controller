package server_files_service

type taskStatus string

const (
	Queued    taskStatus = "queued"
	Executing taskStatus = "executing"
	Completed taskStatus = "completed"
	Failed    taskStatus = "failed"
)

type getTaskStatusResponse struct {
	TaskStatus taskStatus
	Error      error
}

type workerContext struct {
	OperationKind string
	Context       any
	TaskStatus    taskStatus
	TaskId        string
	Error         error
}

type SaveFileRequest struct {
	System string
	S3     struct {
		Endpoint string
		Bucket   string

		ObjectKey string
		FilePath  string
	}
}

type SaveFileResponse struct {
	TaskId string
}

type DeleteFileRequest struct {
}

type DeleteFileResponse struct {
	TaskId string
}

type TransferFileRequest struct {
	System string
	S3     struct {
		Endpoint string
		Bucket   string

		ObjectKey string
		FilePath  string
	}
}

type TransferFileResponse struct {
	TaskId string
}

type PollTaskResponse struct {
	TaskStatus taskStatus
	Error      error
}
