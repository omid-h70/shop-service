package domain

import "context"

type OrderEntity struct {
	OrderId      int64
	VendorId     int64
	CreatedAt    string
	DeliveryTime string
	OrderStatus  string
}

type AddNewOrderRequest struct {
	VendorId int64
}

//type MakeDelayReportRequest struct {
//	OrderId  int64
//	VendorId int64
//}

type DelayReportEntity struct {
	DelayOrderId      int64
	OrderId           int64
	VendorId          int64
	AgentId           int64
	CreatedAt         string
	UpdatedAt         string
	DelayReportStatus string
	ReportCount       int
}

//
//type HandleDelayReportRequest struct {
//	DelayOrderId      int64
//	OrderId           int64
//	VendorId          int64
//	AgentId           int64
//	DelayReportStatus string
//}

type OrderRepository interface {
	AddNewOrder(ctx context.Context, input AddNewOrderRequest) (bool, error)
	InsertDelayReport(ctx context.Context, req DelayReportEntity) (DelayReportEntity, error)
	GetOrderDetails(ctx context.Context, order OrderEntity) (OrderEntity, error)
	GetDelayReportByParams(ctx context.Context, req DelayReportEntity) (DelayReportEntity, error)
	UpdateDelayReport(ctx context.Context, req DelayReportEntity) (bool, error)
	UpdateDelayReportStatus(ctx context.Context, req DelayReportEntity) (bool, error)
}
