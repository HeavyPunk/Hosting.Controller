package server_files_s3_service

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func SaveFileFromS3(req SaveFileFromS3Request) (SaveFileFromS3Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Failed to load default config")
	}
	client := s3.NewFromConfig(cfg, func(opt *s3.Options) {
		opt.BaseEndpoint = &req.Endpoint
	})
	result, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &req.Bucket,
		Key:    &req.ObjectKey,
	})
	if err != nil {
		return SaveFileFromS3Response{}, err
	}
	defer result.Body.Close()

	file, err := os.Create(req.FilePath)
	if err != nil {
		return SaveFileFromS3Response{}, err
	}
	defer file.Close()

	for {
		buff := make([]byte, 1024)
		n, err := result.Body.Read(buff)
		if err == io.EOF {
			_, err = file.Write(buff[:n])
			if err != nil {
				return SaveFileFromS3Response{}, err
			}
			break
		}
		if err != nil {
			return SaveFileFromS3Response{}, err
		}
		_, err = file.Write(buff[:n])
		if err != nil {
			return SaveFileFromS3Response{}, err
		}
	}
	return SaveFileFromS3Response{}, nil
}

func PublishFileToS3(req PublishFileToS3Request) (PublishFileToS3Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Failed to load default config")
	}
	client := s3.NewFromConfig(cfg, func(opt *s3.Options) {
		opt.BaseEndpoint = &req.Endpoint
	})

	file, err := os.Open(req.FilePath)
	if err != nil {
		return PublishFileToS3Response{}, err
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &req.Bucket,
		Key:    &req.ObjectKey,
		Body:   file,
	})
	if err != nil {
		return PublishFileToS3Response{}, err
	}
	return PublishFileToS3Response{
		Success: true,
	}, nil
}
