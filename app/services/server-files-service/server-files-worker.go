package server_files_service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	server_files_s3_service "simple-hosting/controller/app/services/server-files-service/s3"
)

func worker(id int, jobs <-chan workerContext, result chan<- workerContext) {
	for job := range jobs {
		taskCache[job.TaskId].TaskStatus = Executing

		switch job.OperationKind {
		case "save":
			req := job.Context.(SaveFileRequest)
			switch req.System {
			case "S3":
				_, err := server_files_s3_service.SaveFileFromS3(server_files_s3_service.SaveFileFromS3Request{
					Endpoint:  req.S3.Endpoint,
					Bucket:    req.S3.Bucket,
					ObjectKey: req.S3.ObjectKey,
					FilePath:  req.S3.FilePath,
				})
				if err != nil {
					job.Error = err
					taskCache[job.TaskId] = &job
					taskCache[job.TaskId].TaskStatus = Failed
				} else {
					taskCache[job.TaskId] = &job
					taskCache[job.TaskId].TaskStatus = Completed
				}
				result <- job
				break
			}
			break
		case "delete":
			req := job.Context.(DeleteFileRequest)
			err := os.RemoveAll(req.PathToFile)
			if err != nil {
				job.Error = err
				job.TaskStatus = Failed
			}
			taskCache[job.TaskId] = &job
			result <- job
			break
		case "transfer":
			req := job.Context.(TransferFileRequest)
			switch req.System {
			case "S3":
				_, err := server_files_s3_service.PublishFileToS3(server_files_s3_service.PublishFileToS3Request{
					Endpoint:  req.S3.Endpoint,
					Bucket:    req.S3.Bucket,
					ObjectKey: req.S3.ObjectKey,
					FilePath:  req.S3.FilePath,
				})
				if err != nil {
					job.Error = err
					taskCache[job.TaskId] = &job
					taskCache[job.TaskId].TaskStatus = Failed
				} else {
					taskCache[job.TaskId] = &job
					taskCache[job.TaskId].TaskStatus = Completed
				}
				result <- job
				break
			}
			break
		case "create-directory":
			req := job.Context.(CreateDirectoryRequest)
			err := os.MkdirAll(req.PathToDirectory, 0777)
			if err != nil {
				job.Error = err
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Failed
			} else {
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Completed
			}
			result <- job
			break
		case "create-file":
			req := job.Context.(CreateFileRequest)
			// Расшифровываем контент из base64
			content, err := base64.StdEncoding.DecodeString(req.ContentBase64)
			if err != nil {
				job.Error = err
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Failed
				break
			}
			// Создаем файл и записываем в него расшифрованный контент
			err = ioutil.WriteFile(req.PathToFile, content, 0644)
			if err != nil {
				job.Error = err
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Failed
				break
			}
			taskCache[job.TaskId] = &job
			taskCache[job.TaskId].TaskStatus = Completed
			result <- job
		case "list-directory":
			req := job.Context.(ListDirectoryRequest)
			dir := req.PathToDirectory
			var filesList []FileNode
			err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err)
					return nil
				}
				var fileType fileNodeType
				if info.IsDir() {
					fileType = Directory
				} else {
					fileType = File
				}
				filesList = append(filesList, FileNode{
					Path: path,
					Type: fileType,
				})
				return nil
			})
			if err != nil {
				job.Error = err
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Failed
			} else {
				taskCache[job.TaskId] = &job
				taskCache[job.TaskId].TaskStatus = Completed
			}
			result <- job
		default:
			job.Error = errors.New("unknown operation kind " + job.OperationKind)
			result <- job
			break
		}
	}
}

func resultToCacheSaver(results <-chan workerContext) {
	for result := range results {
		taskCache[result.TaskId].Error = result.Error
	}
}
