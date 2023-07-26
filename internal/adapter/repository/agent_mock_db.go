package repository

import (
	"context"
	"github.com/omid-h70/shop-service/internal/core/domain"
)

type AgentRepositoryMockDB struct {
}

func (a AgentRepositoryMockDB) FindAgentLastRecord(ctx context.Context, req domain.AgentSetDelayedOrderRequest) (domain.AgentSetDelayedOrderResponse, error) {
	var resp domain.AgentSetDelayedOrderResponse
	return resp, nil
}

func (a AgentRepositoryMockDB) SetAgentForDelayedOrder(ctx context.Context, req domain.AgentSetDelayedOrderRequest) (bool, error) {
	return false, nil
}

func NewAgentRepositoryMockDB() AgentRepositoryMockDB {
	return AgentRepositoryMockDB{}
}
