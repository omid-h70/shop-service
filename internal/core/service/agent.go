package service

import (
	"context"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"time"
)

type (
	AgentService interface {
		SetAgentToDelayedOrder(context.Context, domain.AgentSetDelayedOrderRequest) (domain.AgentSetDelayedOrderResponse, error)
	}

	AgentServiceImpl struct {
		repo       domain.AgentRepository
		ctxTimeout time.Duration
	}
)

func (a AgentServiceImpl) SetAgentToDelayedOrder(ctx context.Context, req domain.AgentSetDelayedOrderRequest) (domain.AgentSetDelayedOrderResponse, error) {

	resp, err := a.repo.FindAgentLastRecord(ctx, req)
	if err == repository.ErrAgentIsFree {
		var result bool
		result, err = a.repo.SetAgentForDelayedOrder(ctx, req)
		if result && err == nil {
			resp, _ = a.repo.FindAgentLastRecord(ctx, req)
		}
	}
	return resp, err
}

func NewAgentService(repo domain.AgentRepository, t time.Duration) AgentServiceImpl {
	return AgentServiceImpl{
		repo:       repo,
		ctxTimeout: t,
	}
}
