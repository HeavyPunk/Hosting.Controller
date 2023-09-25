package files_controller

import (
	"net/http"
	server_files_service "simple-hosting/controller/app/services/server-files-service"
	"simple-hosting/controller/app/settings"

	errors_tools "simple-hosting/go-commons/tools/errors"

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
		Success:    err == nil,
		Error:      errors_tools.GetErrorStringOrDefault(err, ""),
	})
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
