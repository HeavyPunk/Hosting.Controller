package server_files_service

type taskStatus string

const (
	Queued    taskStatus = "queued"
	Executing taskStatus = "executing"
	Completed taskStatus = "completed"
	Failed    taskStatus = "failed"
)

type fileNodeType string

const (
	Directory fileNodeType = "directory"
	File      fileNodeType = "file"
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
	PathToFile string
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

type CreateDirectoryRequest struct {
	PathToDirectory string
}

type CreateDirectoryResponse struct {
	TaskId string
}

type CreateFileRequest struct {
	PathToFile    string
	ContentBase64 string
}

type CreateFileResponse struct {
	TaskId string
}

type ListDirectoryRequest struct {
	PathToDirectory string
}

type ListDirectoryResponse struct {
	FileNodes []FileNode
}

type FileNode struct {
	Path        string
	Type        fileNodeType
	SizeInBytes int64
	Name        string
	Extension   string
}

type GetFileContentBase64Request struct {
	Path string
}

type GetFileContentBase64Response struct {
	ContentBase64 string
}

type PollTaskResponse struct {
	TaskStatus taskStatus
	Error      error
}
