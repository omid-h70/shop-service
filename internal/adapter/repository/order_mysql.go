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
	ErrDelayReportDoesNotExist = errors.New("Report Doesnt Exist Or Closed")
	ErrOrderDoesNotExist       = errors.New("Order Doesnt Exist")
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

	if err == sql.ErrNoRows {
		return order, ErrOrderDoesNotExist
	}
	return order, err
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
		query1 string = "SELECT delay_report_id, order_id, vendor_id, report_count, delay_report_status, created_at, updated_at " +
			"FROM delay_reports WHERE " +
			"(order_id = ? AND vendor_id = ? AND delay_report_status='OPEN') "

		query2 string = "SELECT delay_report_id, order_id, vendor_id, report_count, delay_report_status, created_at, updated_at " +
			"FROM delay_reports WHERE " +
			"(order_id = ? AND vendor_id = ? AND agent_id = ? AND delay_report_status='OPEN') "

		query string
		resp  domain.DelayReportEntity
	)

	row := &sql.Row{}
	if req.AgentId == 0 {
		query = query1
		row = o.client.QueryRowContext(ctx, query, req.OrderId, req.VendorId)
	} else {
		query = query2
		row = o.client.QueryRowContext(ctx, query, req.OrderId, req.VendorId, req.AgentId)
	}

	err := row.Scan(
		&resp.DelayOrderId,
		&resp.OrderId,
		&resp.VendorId,
		&resp.ReportCount,
		&resp.DelayReportStatus,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return resp, ErrDelayReportDoesNotExist
	case err != nil:
		return resp, err
	}

	if resp.DelayOrderId > 0 {
		err = ErrDelayReportAlreadyExist
	}

	return resp, err
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
	//var (
	//	query string = "UPDATE delay_reports SET delay_report_status = ?, updated_at = now() WHERE (order_id = ? AND agent_id = ?)"
	//)

	var (
		query string = `UPDATE delay_reports ` +
			`INNER JOIN orders ` +
			`ON orders.order_id = delay_reports.order_id ` +
			`SET delay_reports.delay_report_status = ?, orders.delivery_time = now() ` +
			`WHERE (delay_reports.order_id = ? AND delay_reports.vendor_id = ? AND delay_reports.agent_id = ?) `
	)

	_, err := o.client.Exec(query, req.DelayReportStatus, req.OrderId, req.VendorId, req.AgentId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewOrderRepositoryMySqlDB(clientIn *sql.DB) OrderRepositoryMySqlDB {
	return OrderRepositoryMySqlDB{clientIn}
}
