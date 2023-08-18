package utils

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"path"
	"time"

	"go-zero-cloud-disk/core/define"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const PartSize = 5 * 1024 * 1024 // 5MB

func AWSUpload(r *http.Request) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(define.Region))
	if err != nil {
		log.Fatal("Unable to load SDK config", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := s3.NewFromConfig(cfg)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Fatal("Unable to get form file", err)
		return "", err
	}
	defer file.Close()

	key := "cloud-disk/" + GenUUID() + path.Ext(fileHeader.Filename)

	// 1. Initialize the multipart upload
	createResp, err := client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(define.AWSBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal("Failed to create multipart upload", err)
		return "", err
	}

	var completedParts []types.CompletedPart
	partNum := int32(1)

	buffer := make([]byte, PartSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal("Failed to read file part", err)
			return "", err
		}
		if n == 0 {
			break
		}

		// 2. Upload each part
		uploadResp, err := client.UploadPart(ctx, &s3.UploadPartInput{
			Bucket:     aws.String(define.AWSBucket),
			Key:        aws.String(key),
			PartNumber: partNum,
			UploadId:   aws.String(*createResp.UploadId),
			Body:       bytes.NewReader(buffer[:n]),
		})
		if err != nil {
			log.Fatal("Failed to upload part", err)
			return "", err
		}

		completedParts = append(completedParts, types.CompletedPart{
			ETag:       uploadResp.ETag,
			PartNumber: partNum,
		})

		partNum++
	}

	// 3. Complete the multipart upload
	_, err = client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(define.AWSBucket),
		Key:      aws.String(key),
		UploadId: aws.String(*createResp.UploadId),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})
	if err != nil {
		log.Fatal("Failed to complete multipart upload", err)
		return "", err
	}

	return define.AWSBucket + "/" + key, nil
}
