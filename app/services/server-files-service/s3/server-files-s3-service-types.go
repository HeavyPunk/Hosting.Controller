package server_files_s3_service

type SaveFileFromS3Request struct {
	Endpoint string
	Bucket   string

	ObjectKey string
	FilePath  string
}

type SaveFileFromS3Response struct {
}

type PublishFileToS3Request struct {
	Endpoint string
	Bucket   string

	ObjectKey string
	FilePath  string
}

type PublishFileToS3Response struct {
	Success bool
}
