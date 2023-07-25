package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/omid-h70/shop-service/internal/core/service"
	"io"
	"net/http"
	"strconv"
	"time"
)

//var (
//	ErrThereIsNoOpenDelayReports = errors.New("There Is No Open Delay Reports")
//)

type OrderHandler struct {
	service service.OrderService
	mockUrl string
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

type (
	MakeDelayReportRequest struct {
		OrderId  string `json:"order_id"`
		VendorId string `json:"vendor_id"`
	}

	MakeDelayReportResponse struct {
		DelayReportId string `json:"delay_report_id"`
		OrderId       string `json:"order_id"`
		VendorId      string `json:"vendor_id"`
	}

	HandleDelayedOrderRequest struct {
		OrderId  string `json:"order_id"`
		VendorId string `json:"vendor_id"`
		AgentId  string `json:"agent_id"`
	}

	HandleDelayedOrderResponse struct {
		DelayOrderId int64
		HandleDelayedOrderRequest
		ReportCount int
		CreatedAt   string
		UpdatedAt   string
	}
)

type MockData struct {
	ETA int `json:"eta"`
}
type MockResp struct {
	Status bool     `json:"status"`
	Data   MockData `json:"data"`
}

func generateNewOrderTime(mockUrl string) int {
	jsonBody := []byte(``)
	bodyReader := bytes.NewReader(jsonBody)

	mockReq, err := http.NewRequest(http.MethodGet, mockUrl, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return -1
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(mockReq)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return -1
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return -1
	}

	jResp := MockResp{}
	json.Unmarshal(resBody, &jResp)
	return jResp.Data.ETA
}

func (o *OrderHandler) handleMakeDelayReport(w http.ResponseWriter, r *http.Request) {

	var req MakeDelayReportRequest

	if validateJsonRequest(w, r, &req) {

		var (
			domainReq      domain.DelayReportEntity
			vendorParseErr error
			orderParseErr  error
			err            error
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

		domainOrderReq, err = o.service.GetOrderDetails(r.Context(), domainOrderReq)
		if err == repository.ErrOrderDoesNotExist {
			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		}

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
			response.NewSuccess("Delay Time Report Is Invalid - Try After Delivery Time Reached", domainOrderReq, 200).Send(w)
		}
	}
}

func (o *OrderHandler) handleDelayedOrderByAgent(w http.ResponseWriter, r *http.Request) {
	var (
		req HandleDelayedOrderRequest
		//msg    string = "Done"
		result bool
	)

	if validateJsonRequest(w, r, &req) {

		var (
			domainReq                                    domain.DelayReportEntity
			vendorParseErr, orderParseErr, agentParseErr error
		)

		//----------------------------------------------------------

		//Take id as string
		domainReq.OrderId, vendorParseErr = strconv.ParseInt(req.OrderId, 10, 64)
		domainReq.VendorId, orderParseErr = strconv.ParseInt(req.VendorId, 10, 64)
		domainReq.AgentId, agentParseErr = strconv.ParseInt(req.AgentId, 10, 64)

		if orderParseErr != nil && vendorParseErr != nil && agentParseErr != nil {
			response.NewError(ErrInvalidJsonID, http.StatusBadRequest).Send(w)
			return
		}

		//----------------------------------------------------------
		entry := domain.DelayReportEntity{
			VendorId: domainReq.VendorId,
			OrderId:  domainReq.OrderId,
			AgentId:  domainReq.AgentId,
		}
		var err error
		entry, err = o.service.GetDelayReportDetails(r.Context(), entry)
		if err == repository.ErrDelayReportAlreadyExist {
			go func() {
				generateNewOrderTime(o.mockUrl)
				domainReq.DelayReportStatus = "CLOSED"

				_, err = o.service.HandleDelayReport(r.Context(), domainReq)
				if err != nil {
					//msg = err.Error()
				}
			}()

			response.NewSuccess("We're Getting New Update Time For Your Order, Please Check Delay Report Status", true, 200).Send(w)
		} else {
			response.NewSuccess(err.Error(), result, 200).Send(w)
		}
		//----------------------------------------------------------

	}

}

func NewOrderHandler(service service.OrderService, mockUrl string) OrderHandler {
	return OrderHandler{
		service: service,
		mockUrl: mockUrl,
	}
}
