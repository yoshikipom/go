package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	accessKey  string = "root"
	secretKey  string = "miniopass"
	s3Url      string = "http://localhost:9050"
	bucketName string = "mybucket"
	fileName   string = "sample.txt"
)

func buildClient(ctx context.Context) *s3.Client {
	endpoint := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: s3Url,
		}, nil
	})
	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred), config.WithEndpointResolver(endpoint))
	if err != nil {
		log.Fatalln(err)
	}

	return s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})
}

func loadFile(filePath string) (*os.File, func()) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	cleanup := func() { file.Close() }
	return file, cleanup
}

func putFile(ctx context.Context, client s3.Client, file os.File) {
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   &file,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	ctx := context.Background()

	client := buildClient(ctx)

	targetFilePath := fmt.Sprintf("./%s", fileName)
	file, cleanup := loadFile(targetFilePath)
	defer cleanup()

	putFile(ctx, *client, *file)
}
