package service

import (
	"review-service/model"
	"time"
)

type ReviewService struct {
	cachedReviews []model.Review
}

func NewReviewService() *ReviewService {
	cachedReviews := []model.Review{
		{
			ReviewId:  "1",
			ProductId: "1",
			UserId:    "1",
			CreatedAt: time.Now().String(),
			Grade:     5,
		},
		{
			ReviewId:  "2",
			ProductId: "1",
			UserId:    "1",
			CreatedAt: time.Now().String(),
			Grade:     4,
		},
		{
			ReviewId:  "3",
			ProductId: "2",
			UserId:    "1",
			CreatedAt: time.Now().String(),
			Grade:     3,
		},
	}

	return &ReviewService{
		cachedReviews: cachedReviews,
	}
}

func (r ReviewService) Reviews() ([]model.Review, error) {
	return r.cachedReviews, nil
}
