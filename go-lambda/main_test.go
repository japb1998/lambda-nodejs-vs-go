package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestHandler(t *testing.T) {
	// Create a downloader with the session and default options
	sdkConfig, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("personal"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("failed to load default config: %s", err)
	}
	client = s3.NewFromConfig(sdkConfig)

	res, err := handler(context.Background(), events.APIGatewayProxyRequest{})

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	fmt.Println(res)
}
