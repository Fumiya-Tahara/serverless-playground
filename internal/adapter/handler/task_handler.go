package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Fumiya-Tahara/serverless-playground/internal/usecase/task"
	"github.com/aws/aws-lambda-go/events"
)

type TaskHandler struct {
	usecase task.TaskUsecase
}

func NewTaskHandler(u task.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: u}
}

func (h *TaskHandler) CreateTask(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	input := task.CreateTaskInput{
		Title:   body.Title,
		Content: body.Content,
	}

	if err := h.usecase.Create(ctx, input); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func (h *TaskHandler) ListTasks(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	outputs, err := h.usecase.FindAll(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	resBody, _ := json.Marshal(outputs)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(resBody),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	taskID, ok := req.PathParameters["task_id"]
	if !ok {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	input := task.UpdateTaskInput{
		ID:      taskID,
		Title:   body.Title,
		Content: body.Content,
	}

	if err := h.usecase.Update(ctx, input); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	taskID, ok := req.PathParameters["task_id"]
	if !ok {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	if err := h.usecase.Delete(ctx, taskID); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
