package utils

import (
	"context"
	"log"
	"net/http"
	"path"

	"go-zero-cloud-disk/core/define"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AWSUpload(r *http.Request) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(define.Region))
	if err != nil {
		log.Fatal("err: ", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GenUUID() + path.Ext(fileHeader.Filename)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(define.AWSBucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Fatal("failed to upload file, ", err)
		return "", err
	}
	return define.AWSBucket + "/" + key, nil
}
