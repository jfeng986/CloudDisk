package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var region = "us-west-2"

func TestConnectAndUploadToS3(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal("err: ", err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("gobuckettest"),
	})
	if err != nil {
		log.Fatal("err: ", err)
	}
	log.Println("All Objects in bucket")

	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

	// Upload a image to "testbucket" bucket
	filename := "test.png"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file, ", filename, err)
	}
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("gobuckettest"),
		Key:    aws.String("test.png"),
		Body:   f,
	})
	if err != nil {
		log.Fatal("failed to upload file, ", err)
	}
	output, err = client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("gobuckettest"),
	})
	if err != nil {
		log.Fatal("err: ", err)
	}
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}
