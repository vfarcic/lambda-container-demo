package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type NameEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name NameEvent) (string, error) {
	return fmt.Sprintf("%s, please visit https://www.youtube.com/c/TheDevOpsToolkitSeries and subscribe!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
