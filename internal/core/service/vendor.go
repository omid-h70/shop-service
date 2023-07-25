package service

import (
	"context"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"time"
)

type (
	VendorService interface {
		GetAllDelayedOrdersByVendor(context.Context, domain.VendorReportRequest) ([]domain.VendorReportResponse, error)
	}
	//
	VendorServiceImpl struct {
		repo       domain.VendorRepository
		ctxTimeout time.Duration
	}
)

func (v VendorServiceImpl) GetAllDelayedOrdersByVendor(ctx context.Context, req domain.VendorReportRequest) ([]domain.VendorReportResponse, error) {
	return v.repo.GetAllVendorsDelayReports(ctx, req)
}

// NewVendorService do
func NewVendorService(repo domain.VendorRepository, t time.Duration) VendorServiceImpl {
	return VendorServiceImpl{
		repo:       repo,
		ctxTimeout: t,
	}
}
