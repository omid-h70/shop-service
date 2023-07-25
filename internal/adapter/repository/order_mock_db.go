package repository

import (
	"context"
	"github.com/omid-h70/shop-service/internal/core/domain"
)

//
//var (
//	ErrDelayReportAlreadyExist = errors.New("Report Already Exist")
//	ErrDelayReportDoesNotExist = errors.New("Report Doesnt Exist")
//	ErrOrderDoesNotExist       = errors.New("Order Doesnt Exist")
//)

type OrderRepositoryMockDB struct {
}

func (o OrderRepositoryMockDB) AddNewOrder(ctx context.Context, input domain.AddNewOrderRequest) (bool, error) {
	return true, nil
}

func (o OrderRepositoryMockDB) GetOrderDetails(ctx context.Context, order domain.OrderEntity) (domain.OrderEntity, error) {
	return order, nil
}

func (o OrderRepositoryMockDB) InsertDelayReport(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error) {
	return req, nil
}

func (o OrderRepositoryMockDB) GetDelayReportByParams(ctx context.Context, req domain.DelayReportEntity) (domain.DelayReportEntity, error) {
	return domain.DelayReportEntity{}, ErrDelayReportAlreadyExist
}

func (o OrderRepositoryMockDB) UpdateDelayReport(ctx context.Context, req domain.DelayReportEntity) (bool, error) {

	return true, nil
}

func (o OrderRepositoryMockDB) UpdateDelayReportStatus(ctx context.Context, req domain.DelayReportEntity) (bool, error) {
	return true, nil
}

func NewOrderRepositoryMockDB() OrderRepositoryMockDB {
	return OrderRepositoryMockDB{}
}
