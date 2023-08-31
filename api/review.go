package api

import (
	"net/http"
	"review-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ReviewAPI struct {
	reviewService *service.ReviewService
	validator     *validator.Validate
}

func NewReviewAPI(reviewService *service.ReviewService, validator *validator.Validate) *ReviewAPI {
	return &ReviewAPI{
		reviewService: reviewService,
		validator:     validator,
	}
}

func (r *ReviewAPI) Register(server *echo.Echo) {
	server.GET("/v1/reviews", r.listAll)
	server.POST("/v1/reviews", r.createReview)
}

func (r *ReviewAPI) listAll(c echo.Context) error {
	reviews, err := r.reviewService.Reviews()
	if err != nil {

	}

	return c.JSON(http.StatusOK, reviews)
}

func (r *ReviewAPI) createReview(c echo.Context) error {
	payload := ReviewPayload{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.validator.Struct(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload)
}
