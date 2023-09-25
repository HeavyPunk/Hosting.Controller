package server_files_service

import (
	"errors"
	"simple-hosting/controller/app/settings"

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
