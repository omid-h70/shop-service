package service

import (
	"context"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"time"
)

type (
	OrderService interface {
		AddNewOrder(ctx context.Context, req domain.AddNewOrderRequest) error
		GetOrderDetails(ctx context.Context, order domain.OrderEntity) (domain.OrderEntity, error)
		GetDelayReportDetails(ctx context.Context, order domain.DelayReportEntity) (domain.DelayReportEntity, error)
		AddOrUpdateDelayReport(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error)
		HandleDelayReport(ctx context.Context, req domain.DelayReportEntity) (bool, error)
	}

	OrderServiceImpl struct {
		repo       domain.OrderRepository
		ctxTimeout time.Duration
	}
)

func (t OrderServiceImpl) AddNewOrder(ctx context.Context, req domain.AddNewOrderRequest) error {
	_, err := t.repo.AddNewOrder(ctx, req)
	return err
}

func (t OrderServiceImpl) GetDelayReportDetails(ctx context.Context, order domain.DelayReportEntity) (domain.DelayReportEntity, error) {
	return t.repo.GetDelayReportByParams(ctx, order)
}

func (t OrderServiceImpl) GetOrderDetails(ctx context.Context, order domain.OrderEntity) (domain.OrderEntity, error) {
	return t.repo.GetOrderDetails(ctx, order)
}

func (t OrderServiceImpl) AddOrUpdateDelayReport(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error) {

	resp, err := t.repo.GetDelayReportByParams(ctx, req)
	if err == repository.ErrDelayReportAlreadyExist {

		result, err := t.repo.UpdateDelayReport(ctx, req)
		if result {
			resp.UpdatedAt = time.Now().String()
		}
		return resp, err

	} else if err == nil || err == repository.ErrDelayReportDoesNotExist {

		resp, err = t.repo.InsertDelayReport(ctx, req)
		if err == nil {
			resp, err = t.repo.GetDelayReportByParams(ctx, req)
		}
		return resp, nil
	}
	return resp, err
}

func (t OrderServiceImpl) HandleDelayReport(ctx context.Context, req domain.DelayReportEntity) (bool, error) {
	return t.repo.UpdateDelayReportStatus(ctx, req)
}

// NewOrderService do
func NewOrderService(repo domain.OrderRepository, t time.Duration) OrderServiceImpl {
	return OrderServiceImpl{
		repo:       repo,
		ctxTimeout: t,
	}
}
