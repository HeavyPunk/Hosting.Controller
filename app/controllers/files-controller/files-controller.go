package files_controller

import (
	"net/http"
	server_files_service "simple-hosting/controller/app/services/server-files-service"
	"simple-hosting/controller/app/settings"

	errors_tools "simple-hosting/go-commons/tools/errors"
	tools_sequence "simple-hosting/go-commons/tools/sequence"

	"github.com/gin-gonic/gin"
)

func PushFileS3(c *gin.Context) {
	var request PushFileS3Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	config, err := settings.GetSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, PushFileS3Response{
			Success: false,
			Error:   errors_tools.GetErrorStringOrDefault(err, ""),
		})
		return
	}
	resp, err := server_files_service.TransferFile(server_files_service.TransferFileRequest{
		System: "S3",
		S3: struct {
			Endpoint  string
			Bucket    string
			ObjectKey string
			FilePath  string
		}{
			Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
			Bucket:    request.S3Bucket,
			ObjectKey: request.S3FileName,
			FilePath:  request.LocalFilePath,
		},
	})
	c.JSON(http.StatusOK, PushFileS3Response{
		TaskId:  resp.TaskId,
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func PollTask(c *gin.Context) {
	var request PollTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_files_service.PollTask(request.TaskId)
	c.JSON(http.StatusOK, PollTaskResponse{
		TaskStatus: string(resp.TaskStatus),
		TaskError:  errors_tools.GetErrorStringOrDefault(resp.Error, ""),
		Success:    err == nil,
		Error:      errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func AcceptTask(c *gin.Context) {
	var request AcceptTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	server_files_service.ConfirmExecution(request.TaskId)
}

func GetFileS3(c *gin.Context) {
	var request GetFileS3Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	config, err := settings.GetSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, PushFileS3Response{
			Success: false,
			Error:   errors_tools.GetErrorStringOrDefault(err, ""),
		})
		return
	}

	resp, err := server_files_service.SaveFile(server_files_service.SaveFileRequest{
		System: "S3",
		S3: struct {
			Endpoint  string
			Bucket    string
			ObjectKey string
			FilePath  string
		}{
			Endpoint:  config.App.Services.ServerFiles.S3.Endpoint,
			Bucket:    request.S3Bucket,
			ObjectKey: request.S3FileName,
			FilePath:  request.LocalFilePath,
		},
	})
	c.JSON(http.StatusOK, GetFileS3Response{
		TaskId:  resp.TaskId,
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func RemoveFile(c *gin.Context) {
	var request RemoveFileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_files_service.DeleteFile(server_files_service.DeleteFileRequest{
		PathToFile: request.PathToFile,
	})
	c.JSON(http.StatusOK, RemoveFileResponse{
		TaskId:  resp.TaskId,
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func CreateFile(c *gin.Context) {
	var request CreateFileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_files_service.CreateFile(server_files_service.CreateFileRequest{
		PathToFile:    request.PathToFile,
		ContentBase64: request.ContentBase64,
	})

	c.JSON(http.StatusOK, CreateFileResponse{
		TaskId:  resp.TaskId,
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func CreateDirectory(c *gin.Context) {
	var request CreateDirectoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_files_service.CreateDirectory(server_files_service.CreateDirectoryRequest{
		PathToDirectory: request.PathToDirectory,
	})

	c.JSON(http.StatusOK, CreateDirectoryResponse{
		TaskId:  resp.TaskId,
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func ListDirectory(c *gin.Context) {
	var request ListDirectoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}
	resp, err := server_files_service.ListDirectory(server_files_service.ListDirectoryRequest{
		PathToDirectory: request.PathToDirectory,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ListDirectoryResponse{
			Success: err == nil,
			Error:   errors_tools.GetErrorStringOrDefault(err, ""),
		})
	}

	c.JSON(http.StatusOK, ListDirectoryResponse{
		Files: tools_sequence.Mapper(resp.FileNodes, func(node server_files_service.FileNode) FileNode {
			return FileNode{
				Path:      node.Path,
				Type:      string(node.Type),
				Size:      node.SizeInBytes,
				Name:      node.Name,
				Extension: node.Extension,
			}
		}),
		Success: err == nil,
		Error:   errors_tools.GetErrorStringOrDefault(err, ""),
	})
}

func GetFileContentBase64(c *gin.Context) {
	var request GetFileContentBase64Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	resp, err := server_files_service.GetFileContentBase64(server_files_service.GetFileContentBase64Request{
		Path: request.Path,
	})

	c.JSON(http.StatusOK, GetFileContentBase64Response{
		ContentBase64: resp.ContentBase64,
		Success:       err == nil,
		Error:         errors_tools.GetErrorStringOrDefault(err, ""),
	})
}
