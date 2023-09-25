package files_controller

type PollTaskRequest struct {
	TaskId string `json:"task-id"`
}

type PollTaskResponse struct {
	TaskStatus string `json:"task-status"`
	Success    bool   `json:"success"`
	Error      string `json:"error"`
}

type PushFileS3Request struct {
	S3Bucket      string `json:"s3-bucket"`
	S3FileName    string `json:"s3-file-name"`
	LocalFilePath string `json:"local-file-path"`
}

type PushFileS3Response struct {
	TaskId  string `json:"task-id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type GetFileS3Request struct {
	S3Bucket      string `json:"s3-bucket"`
	S3FileName    string `json:"s3-file-name"`
	LocalFilePath string `json:"local-file-path"`
}

type GetFileS3Response struct {
	TaskId  string `json:"task-id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
