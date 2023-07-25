package repository

import (
	"context"
	"database/sql"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/pkg/errors"
)

var (
	ErrVendorRepoDBInvalidExp = errors.New("Error While Fetching Data Rows")
	ErrVendorQueryContext     = errors.New("Error While Query Execution")
)

type VendorRepositoryMySqlDB struct {
	client *sql.DB
}

func (v *VendorRepositoryMySqlDB) GetAllVendorsDelayReports(ctx context.Context, input domain.VendorReportRequest) ([]domain.VendorReportResponse, error) {

	/*
		SELECT
		delay_reports.vendor_id,
		-- Sum(Time_to_sec(Date_format(orders.delivery_time, '%h:%i:%s')) + Time_to_sec(Date_format(delay_reports.updated_at, '%h:%i:%s'))) As delay_time
		-- SUM(timestampdiff(minute, delay_reports.updated_at, orders.delivery_time)) as delay_time
		timestampdiff(minute, delay_reports.updated_at, orders.delivery_time) as delay_time
		FROM delay_reports
		INNER JOIN orders ON delay_reports.vendor_id = orders.vendor_id
		WHERE DATE(delay_reports.created_at) < (NOW() - INTERVAL 7 DAY)
		GROUP BY delay_reports.vendor_id
	*/

	var (
		query string = `SELECT ` +
			`delay_reports.vendor_id, ` +
			//-- Sum(Time_to_sec(Date_format(orders.delivery_time, '%h:%i:%s')) + Time_to_sec(Date_format(delay_reports.updated_at, '%h:%i:%s'))) As delay_time
			`SUM(timestampdiff(minute, delay_reports.updated_at, orders.delivery_time)) as delay_time ` +
			//`timestampdiff(minute, delay_reports.updated_at, orders.delivery_time) as delay_time ` +
			`FROM delay_reports ` +
			`INNER JOIN orders ON delay_reports.vendor_id = orders.vendor_id ` +
			`WHERE DATE(delay_reports.created_at) > (NOW() - INTERVAL 7 DAY) ` +
			`GROUP BY delay_reports.vendor_id `
	)

	rows, err := v.client.QueryContext(ctx, query)
	if err != nil {
		return []domain.VendorReportResponse{}, ErrVendorQueryContext
	}
	//
	var records = make([]domain.VendorReportResponse, 0)
	for rows.Next() {
		record := domain.VendorReportResponse{}
		//
		if err = rows.Scan(&record.ID, &record.DelayTime); err != nil {
			return []domain.VendorReportResponse{}, ErrVendorRepoDBInvalidExp
		}
		records = append(records, record)
	}
	//
	return records, nil
}

func NewVendorRepositoryMySqlDB(clientIn *sql.DB) *VendorRepositoryMySqlDB {
	return &VendorRepositoryMySqlDB{clientIn}
}
