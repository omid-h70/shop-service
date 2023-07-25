package service

import (
	"github.com/omid-h70/shop-service/internal/core/domain"
	"time"
)

type (
	TripService interface {
	}

	TripServiceImpl struct {
		repo       domain.TripRepository
		ctxTimeout time.Duration
	}
)
