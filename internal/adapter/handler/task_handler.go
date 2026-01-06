package handler

import (
	"net/http"

	"github.com/Fumiya-Tahara/serverless-playground/internal/usecase/task"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	usecase task.TaskUsecase
}

func NewTaskHandler(u task.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: u}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	type request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	input := task.CreateTaskInput{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.usecase.Create(c.Request().Context(), input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (h *TaskHandler) ListTasks(c echo.Context) error {
	outputs, err := h.usecase.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to fetch tasks",
		})
	}

	return c.JSON(http.StatusOK, outputs)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskID := c.Param("task_id")

	type request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	input := task.UpdateTaskInput{
		ID:      taskID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.usecase.Update(c.Request().Context(), input); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("task_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "task_id is required",
		})
	}

	if err := h.usecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete task",
		})
	}

	return c.NoContent(http.StatusOK)
}
