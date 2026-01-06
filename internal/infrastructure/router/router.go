package router

import (
	"github.com/Fumiya-Tahara/serverless-playground/internal/adapter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(h *handler.TaskHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	g := e.Group("/api/v1")
	{
		g.POST("/tasks", h.CreateTask)
		g.GET("/tasks", h.ListTasks)
		g.PATCH("/tasks/:task_id", h.UpdateTask)
		g.DELETE("/tasks/:task_id", h.DeleteTask)
	}

	return e
}
