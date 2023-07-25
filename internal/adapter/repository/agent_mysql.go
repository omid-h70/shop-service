package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/omid-h70/shop-service/internal/core/domain"
)

var (
	ErrAgentISAlreadyBusy = errors.New("Agent is Already Busy")
	ErrAgentIsFree        = errors.New("Agent is Free")
	ErrNoAffectedRows     = errors.New("No Affected Rows")
)

type AgentRepositoryMySqlDB struct {
	client *sql.DB
}

func (o AgentRepositoryMySqlDB) FindAgentLastRecord(ctx context.Context, req domain.AgentSetDelayedOrderRequest) (domain.AgentSetDelayedOrderResponse, error) {
	var (
		query string = "SELECT delay_report_id, order_id, vendor_id, agent_id, report_count, created_at, updated_at FROM delay_reports WHERE (agent_id= ?)"
		resp  domain.AgentSetDelayedOrderResponse
	)

	err := o.client.QueryRowContext(ctx, query, req.AgentId).Scan(
		&resp.DelayOrderId,
		&resp.OrderId,
		&resp.VendorId,
		&resp.AgentId,
		&resp.ReportCount,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return resp, ErrAgentIsFree
	case err != nil:
		return resp, err
	}

	return resp, ErrAgentISAlreadyBusy

}

func (o AgentRepositoryMySqlDB) SetAgentForDelayedOrder(ctx context.Context, req domain.AgentSetDelayedOrderRequest) (bool, error) {
	var (
		query string = "UPDATE delay_reports SET agent_id = ? WHERE agent_id IS NULL ORDER BY created_at LIMIT 1"
	)

	res, err := o.client.Exec(query, req.AgentId)
	if err != nil {
		return false, err
	} else {
		if count, err := res.RowsAffected(); count == 0 && err == nil {
			return false, ErrNoAffectedRows
		}
	}
	return true, nil
}

func NewAgentRepositoryMySqlDB(clientIn *sql.DB) AgentRepositoryMySqlDB {
	return AgentRepositoryMySqlDB{clientIn}
}
