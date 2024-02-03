package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

func main() {
	// Create a downloader with the session and default options
	sdkConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load default config: %s", err)
	}
	client = s3.NewFromConfig(sdkConfig)
	lambda.Start(handler)
}

func handler(ctx context.Context, e events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	input := s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String("hello.txt"),
	}
	out, err := client.GetObject(ctx, &input)

	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer([]byte{})
	buff.ReadFrom(out.Body)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: buff.String(),
	}, nil
}
