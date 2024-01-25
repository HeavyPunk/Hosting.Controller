package files_controller

type PollTaskRequest struct {
	TaskId string `json:"task-id"`
}

type PollTaskResponse struct {
	TaskStatus string `json:"task-status"`
	TaskError  string `json:"task-error"`
	Success    bool   `json:"success"`
	Error      string `json:"error"`
}

type AcceptTaskRequest struct {
	TaskId string `json:"task-id"`
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

type RemoveFileRequest struct {
	PathToFile string `json:"path"`
}

type RemoveFileResponse struct {
	TaskId  string `json:"task-id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type CreateFileRequest struct {
	PathToFile    string `json:"path"`
	ContentBase64 string `json:"content-base64"`
}

type CreateFileResponse struct {
	TaskId  string `json:"task-id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type CreateDirectoryRequest struct {
	PathToDirectory string `json:"path"`
}

type CreateDirectoryResponse struct {
	TaskId  string `json:"task-id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ListDirectoryRequest struct {
	PathToDirectory string `json:"path"`
}

type FileNode struct {
	Path      string `json:"path"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type ListDirectoryResponse struct {
	Files   []FileNode `json:"files"`
	Success bool       `json:"success"`
	Error   string     `json:"error"`
}

type GetFileContentBase64Request struct {
	Path string `json:"path"`
}

type GetFileContentBase64Response struct {
	ContentBase64 string `json:"content-base64"`
	Success       bool   `json:"success"`
	Error         string `json:"error"`
}
