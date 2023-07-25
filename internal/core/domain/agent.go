package domain

import "context"

type (
	AgentSetDelayedOrderRequest struct {
		OrderId  int64
		VendorId int64
		AgentId  int64
	}

	AgentSetDelayedOrderResponse struct {
		DelayOrderId int64
		AgentSetDelayedOrderRequest
		ReportCount int
		CreatedAt   string
		UpdatedAt   string
	}

	HandleDelayedOrderRequest struct {
		OrderId  int64
		VendorId int64
		AgentId  int64
	}

	HandleDelayedOrderResponse struct {
		DelayOrderId int64
		AgentSetDelayedOrderRequest
		ReportCount int
		CreatedAt   string
		UpdatedAt   string
	}
)

type (
	AgentRepository interface {
		FindAgentLastRecord(ctx context.Context, req AgentSetDelayedOrderRequest) (AgentSetDelayedOrderResponse, error)
		SetAgentForDelayedOrder(ctx context.Context, req AgentSetDelayedOrderRequest) (bool, error)
	}
)
