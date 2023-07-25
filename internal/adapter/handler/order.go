package handler

import (
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultReportMinutesTime = 10
)

type OrderHandler struct {
	service service.OrderService
}

type AddNewOrderRequest struct {
	VendorId int64 `json:"vendor_id"`
}

func (o *OrderHandler) HandleAddNewOrder(w http.ResponseWriter, r *http.Request) {

	var req AddNewOrderRequest

	if validateJsonRequest(w, r, &req) {
		domainReq := domain.AddNewOrderRequest{
			VendorId: req.VendorId,
		}

		err := o.service.AddNewOrder(r.Context(), domainReq)
		if err != nil {
			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		}
		response.NewSuccess("", "", 200).Send(w)
	}
}

type MakeDelayReportRequest struct {
	OrderId  string `json:"order_id"`
	VendorId string `json:"vendor_id"`
}

type MakeDelayReportResponse struct {
	DelayReportId string `json:"delay_report_id"`
	OrderId       string `json:"order_id"`
	VendorId      string `json:"vendor_id"`
}

func (o *OrderHandler) handleMakeDelayReport(w http.ResponseWriter, r *http.Request) {

	var req MakeDelayReportRequest

	if validateJsonRequest(w, r, &req) {

		var (
			domainReq      domain.DelayReportEntity
			vendorParseErr error
			orderParseErr  error
			//err            error
		)

		//Take id as string
		domainReq.OrderId, vendorParseErr = strconv.ParseInt(req.OrderId, 10, 64)
		domainReq.VendorId, orderParseErr = strconv.ParseInt(req.VendorId, 10, 64)

		if orderParseErr != nil && vendorParseErr != nil {
			response.NewError(ErrInvalidJsonID, http.StatusBadRequest).Send(w)
			return
		}

		domainOrderReq := domain.OrderEntity{
			OrderId:  domainReq.OrderId,
			VendorId: domainReq.VendorId,
		}

		domainOrderReq, _ = o.service.GetOrderDetails(r.Context(), domainOrderReq)
		t1, _ := time.Parse("2006-01-02 03:04:05", domainOrderReq.DeliveryTime)

		layout := "2006-01-02 03:04:05 PM"
		now := time.Now().Format(layout)
		t_now, _ := time.Parse(layout, now)

		//fmt.Println(t.Unix(), time.Now().Unix())

		if t1.Before(t_now) {
			resp, err := o.service.AddOrUpdateDelayReport(r.Context(), domainReq)
			if err != nil {
				response.NewError(err, http.StatusBadRequest).Send(w)
				return
			}
			response.NewSuccess("Done", resp, 200).Send(w)
		} else {
			response.NewSuccess("Delay Report Is Invalid", domainOrderReq, 200).Send(w)
		}
	}
}

func (o *OrderHandler) handleCloseDelayedOrderByAgent(w http.ResponseWriter, r *http.Request) {
	var (
		req    SetAgentForDelayedOrderRequest
		msg    string = "Done"
		err    error
		result bool
	)

	if validateJsonRequest(w, r, &req) {

		var (
			domainReq                                    domain.DelayReportEntity
			vendorParseErr, orderParseErr, agentParseErr error
		)
		//Take id as string
		//domainReq.OrderId, vendorParseErr = strconv.ParseInt(req.OrderId, 10, 64)
		//domainReq.VendorId, orderParseErr = strconv.ParseInt(req.VendorId, 10, 64)
		domainReq.AgentId, agentParseErr = strconv.ParseInt(req.AgentId, 10, 64)

		if orderParseErr != nil && vendorParseErr != nil && agentParseErr != nil {
			response.NewError(ErrInvalidJsonID, http.StatusBadRequest).Send(w)
			return
		}

		result, err = o.service.HandleDelayReport(r.Context(), domainReq)

		if err != nil {
			msg = err.Error()
		}
	}
	response.NewSuccess(msg, result, 200).Send(w)
}
