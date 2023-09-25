package server_files_service

import (
	"errors"
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
					taskCache[job.TaskId].TaskStatus = Failed
				} else {
					taskCache[job.TaskId].TaskStatus = Completed
				}
				result <- job
				break
			}
			break
		case "delete":
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
					taskCache[job.TaskId].TaskStatus = Failed
				} else {
					taskCache[job.TaskId].TaskStatus = Completed
				}
				result <- job
				break
			}
			break
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
