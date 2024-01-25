package server_files_service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"simple-hosting/controller/app/settings"
	"strings"

	"github.com/docker/distribution/uuid"
)

var taskCache = make(map[string]*workerContext)
var jobs chan workerContext
var results chan workerContext

func pushToQueue(operationKind string, context any) (string, error) {
	taskId := uuid.Generate()
	wContext := workerContext{
		OperationKind: operationKind,
		Context:       context,
		TaskStatus:    Queued,
		TaskId:        taskId.String(),
	}
	taskCache[taskId.String()] = &wContext
	jobs <- wContext
	return taskId.String(), nil
}

func getTaskInfo(taskId string) (getTaskStatusResponse, error) {
	res, ok := taskCache[taskId]
	if !ok {
		return getTaskStatusResponse{}, errors.New("task " + taskId + " not found")
	}
	return getTaskStatusResponse{
		TaskStatus: res.TaskStatus,
		Error:      res.Error,
	}, nil
}

func removeTaskFromCache(taskId string) {
	delete(taskCache, taskId)
}

func Init(settings settings.ServiceSettings) {
	jobs = make(chan workerContext, settings.App.Services.ServerFiles.WorkerPull.QueueSize)
	results = make(chan workerContext, settings.App.Services.ServerFiles.WorkerPull.QueueSize)
	go func() {
		go resultToCacheSaver(results)
		for w := 1; w < settings.App.Services.ServerFiles.WorkerPull.WorkerCount; w++ {
			go worker(w, jobs, results)
		}
	}()
}

func PollTask(taskId string) (PollTaskResponse, error) {
	taskInfo, err := getTaskInfo(taskId)
	return PollTaskResponse{
		TaskStatus: taskInfo.TaskStatus,
		Error:      taskInfo.Error,
	}, err
}

func ConfirmExecution(taskId string) {
	removeTaskFromCache(taskId)
}

func SaveFile(req SaveFileRequest) (SaveFileResponse, error) {
	taskId, err := pushToQueue("save", req)
	return SaveFileResponse{taskId}, err
}

func DeleteFile(req DeleteFileRequest) (DeleteFileResponse, error) {
	taskId, err := pushToQueue("delete", req)
	return DeleteFileResponse{taskId}, err
}

func TransferFile(req TransferFileRequest) (TransferFileResponse, error) {
	taskId, err := pushToQueue("transfer", req)
	return TransferFileResponse{taskId}, err
}

func CreateDirectory(req CreateDirectoryRequest) (CreateDirectoryResponse, error) {
	taskId, err := pushToQueue("create-directory", req)
	return CreateDirectoryResponse{taskId}, err
}

func ListDirectory(req ListDirectoryRequest) (ListDirectoryResponse, error) {
	dir := req.PathToDirectory
	var filesList []FileNode
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		var fileType fileNodeType
		var fileExtension string
		if info.IsDir() {
			fileType = Directory
			fileExtension = ""
		} else {
			fileType = File
			extension := strings.Split(info.Name(), ".")
			fileExtension = extension[len(extension)-1]
		}
		filesList = append(filesList, FileNode{
			Path:        strings.TrimLeft(strings.Replace(path, dir, "", -1), "/"),
			Type:        fileType,
			SizeInBytes: info.Size(),
			Name:        info.Name(),
			Extension:   fileExtension,
		})
		return nil
	})
	if err != nil {
		return ListDirectoryResponse{}, err
	}
	return ListDirectoryResponse{
		FileNodes: filesList,
	}, nil
}

func CreateFile(req CreateFileRequest) (CreateFileResponse, error) {
	taskId, err := pushToQueue("create-file", req)
	return CreateFileResponse{taskId}, err
}

func GetFileContentBase64(req GetFileContentBase64Request) (GetFileContentBase64Response, error) {
	fileContent, err := ioutil.ReadFile(req.Path)
	if err != nil {
		return GetFileContentBase64Response{}, err
	}
	base64Content := base64.StdEncoding.EncodeToString(fileContent)
	return GetFileContentBase64Response{base64Content}, nil
}
