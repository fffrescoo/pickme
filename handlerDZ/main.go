package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func PostHandler(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON, ты че, дебил?"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task is empty, долбоёб"})
	}

	task = req.Task
	return c.JSON(http.StatusOK, map[string]string{"message": "Task saved, красава"})
}
func GetHandler(c echo.Context) error {
	if task == "" {
		return c.JSON(http.StatusOK, map[string]string{"message": "hello, nobody"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "hello, " + task})
}
func main() {
	e := echo.New()

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
