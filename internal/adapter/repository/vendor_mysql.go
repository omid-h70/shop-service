package repository

import (
	"context"
	"database/sql"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/pkg/errors"
)

var (
	ErrTransRepoDBInvalidTransactionExp             = errors.New("invalid NULL Transaction Exception")
	ErrTransRepoDBUpdatingBalanceFailed             = errors.New("Error while updating Balance")
	ErrTransRepoDBWhileFetchingAccountInfo          = errors.New("Error while Fetching Account Info")
	ErrTransRepoDBWhileInsertingTransaction         = errors.New("Error In Transaction")
	ErrTransRepoDBCardIdToCardFromHasTheSameAccount = errors.New("CardFrom and CardTo Has The Same Account")
)

type VendorRepositoryMySqlDB struct {
	client *sql.DB
}

func (v *VendorRepositoryMySqlDB) GetDelayedOrdersByVendor(ctx context.Context, input domain.VendorReportRequest) ([]domain.VendorReportResponse, error) {
	return []domain.VendorReportResponse{}, nil
}

func (v *VendorRepositoryMySqlDB) FindMostActiveCustomersWithinTime(ctx context.Context, count int, time int) ([]domain.VendorReportResponse, error) {

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

	//var (
	//	query string = fmt.Sprintf(`with top_txn_analysis
	//		as
	//		(
	//			select ctmr_txn_cnt.customer_id, txn.*, row_number() over (partition by ctmr_txn_cnt.customer_id order by txn.transaction_date desc) as row_num from transaction txn
	//			left join card crd on txn.card_id_from = crd.card_id or txn.card_id_to = crd.card_id
	//			left join account acnt on crd.account_id = acnt.account_id
	//			right join
	//			(
	//				select ctmr_txn.customer_id from
	//				(
	//					select ctmr.customer_id, txn.* from transaction txn
	//					left join card crd on txn.card_id_from = crd.card_id or txn.card_id_to = crd.card_id
	//					left join account acnt on crd.account_id = acnt.account_id
	//					left join customer ctmr on acnt.customer_id = ctmr.customer_id
	//					WHERE txn.transaction_date >= (NOW() - INTERVAL %d minute)
	//				) as ctmr_txn group by ctmr_txn.customer_id order by count(0) desc limit %d
	//			) as ctmr_txn_cnt on ctmr_txn_cnt.customer_id = acnt.customer_id
	//		) select * from top_txn_analysis where top_txn_analysis.row_num <= 10;
	//		`, time, count)
	//)
	//
	//rows, err := c.client.QueryContext(ctx, query)
	//if err != nil {
	//	return []domain.CustomerReportOut{}, errors.Wrap(err, "error listing accounts")
	//}
	//
	//var records = make([]domain.CustomerReportOut, 0)
	//for rows.Next() {
	//	record := domain.CustomerReportOut{}
	//
	//	if err = rows.Scan(&record.CustomerID, &record.TransactionId, &record.CardIdFrom, &record.CardIdTo, &record.Amount, &record.TransactionType, &record.TransactionTime, &record.Index); err != nil {
	//		return []domain.CustomerReportOut{}, errors.Wrap(err, "error listing accounts")
	//	}
	//	records = append(records, record)
	//}
	//
	return []domain.VendorReportResponse{}, nil
}

func NewVendorRepositoryMySqlDB(clientIn *sql.DB) *VendorRepositoryMySqlDB {
	return &VendorRepositoryMySqlDB{clientIn}
}
