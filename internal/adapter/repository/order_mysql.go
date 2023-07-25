package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/pkg/errors"
)

var (
	ErrDelayReportAlreadyExist = errors.New("Report Already Exist")
	ErrDelayReportDoesNotExist = errors.New("Report Doesnt Exist")
)

type OrderRepositoryMySqlDB struct {
	client *sql.DB
}

func (o OrderRepositoryMySqlDB) AddNewOrder(ctx context.Context, input domain.AddNewOrderRequest) (bool, error) {

	var (
		query string = fmt.Sprintf("INSERT INTO orders (vendor_id) VALUES (?)")
	)

	res, err := o.client.Exec(query, input.VendorId)
	if err != nil {
		return false, err
	}
	lastId, err := res.LastInsertId()
	fmt.Println(lastId)

	return true, nil
}

func (o OrderRepositoryMySqlDB) GetOrderDetails(ctx context.Context, order domain.OrderEntity) (domain.OrderEntity, error) {
	var (
		query string = "SELECT * FROM orders WHERE (order_id = ? AND vendor_id = ?)"
	)

	err := o.client.QueryRowContext(ctx, query, order.OrderId, order.VendorId).Scan(
		&order.OrderId,
		&order.VendorId,
		&order.CreatedAt,
		&order.DeliveryTime,
		&order.OrderStatus,
	)

	switch {
	case err == sql.ErrNoRows:
		return order, ErrDelayReportDoesNotExist
	case err != nil:
		return order, err
	}

	return order, ErrDelayReportAlreadyExist

}

func (o OrderRepositoryMySqlDB) InsertDelayReport(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error) {
	var (
		query string = `INSERT INTO delay_reports (order_id, vendor_id) VALUES (?, ?)`
		resp  domain.DelayReportEntity
	)
	res, err := o.client.Exec(query, req.OrderId, req.VendorId)
	if err != nil {
		return resp, err
	}

	resp.DelayOrderId, _ = res.LastInsertId()
	resp.OrderId, resp.VendorId = req.OrderId, req.VendorId
	//fmt.Println(res.LastInsertId())
	return resp, nil
}

func (o OrderRepositoryMySqlDB) GetDelayReportByParams(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error) {
	var (
		query string = "SELECT delay_report_id, order_id, vendor_id, report_count, created_at, updated_at FROM delay_reports WHERE (order_id = ? AND vendor_id = ?)"
		resp  domain.DelayReportEntity
	)

	err := o.client.QueryRowContext(ctx, query, req.OrderId, req.VendorId).Scan(
		&resp.DelayOrderId,
		&resp.OrderId,
		&resp.VendorId,
		&resp.ReportCount,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return resp, ErrDelayReportDoesNotExist
	case err != nil:
		return resp, err
	}

	return resp, ErrDelayReportAlreadyExist
}

func (o OrderRepositoryMySqlDB) UpdateDelayReport(ctx context.Context, req domain.DelayReportEntity) (bool, error) {
	var (
		query string = "UPDATE delay_reports SET report_count = report_count + 1, updated_at = now() WHERE (order_id = ? AND vendor_id = ?)"
	)

	_, err := o.client.Exec(query, req.OrderId, req.VendorId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o OrderRepositoryMySqlDB) UpdateDelayReportStatus(ctx context.Context, req domain.DelayReportEntity) (bool, error) {
	var (
		query string = "UPDATE delay_reports SET delay_report_status = ?, updated_at = now() WHERE (order_id = ? AND agent_id = ?)"
	)

	_, err := o.client.Exec(query, req.DelayReportStatus, req.OrderId, req.AgentId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewOrderRepositoryMySqlDB(clientIn *sql.DB) OrderRepositoryMySqlDB {
	return OrderRepositoryMySqlDB{clientIn}
}
