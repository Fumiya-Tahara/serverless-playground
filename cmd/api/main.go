// cmd/api/main.go
package main

import (
	"context"

	"github.com/Fumiya-Tahara/serverless-playground/internal/adapter/handler"
	"github.com/Fumiya-Tahara/serverless-playground/internal/adapter/persistence/stub"
	"github.com/Fumiya-Tahara/serverless-playground/internal/infrastructure/router"
	"github.com/Fumiya-Tahara/serverless-playground/internal/usecase/task"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	repo := stub.NewMemoryTaskRepository()
	interactor := task.NewTaskInteractor(repo)
	taskHandler := handler.NewTaskHandler(interactor)

	e := router.NewRouter(taskHandler)

	echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
