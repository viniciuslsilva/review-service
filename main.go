package main

import (
	"net/http"
	"review-service/api"
	"review-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	reviewService := service.NewReviewService()
	validator := validator.New()
	reviewAPI := api.NewReviewAPI(reviewService, validator)
	reviewAPI.Register(e)

	// Start server
	e.Logger.Fatal(e.Start("localhost:8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	// return c.String(http.StatusOK, "Hello, World!")
}
