package handler

import (
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
	"strconv"
)

type AgentHandler struct {
	service service.AgentService
}

type SetAgentForDelayedOrderRequest struct {
	OrderId  string `json:"order_id"`
	VendorId string `json:"vendor_id"`
	AgentId  string `json:"agent_id"`
}

type SetAgentForDelayedOrderResponse struct {
	SetAgentForDelayedOrderRequest
	ReportCount string `json:"report_count"`
	CreateAt    string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CloseDelayedOrderRequest struct {
	OrderId  string `json:"order_id"`
	VendorId string `json:"vendor_id"`
	AgentId  string `json:"agent_id"`
}

func (a *AgentHandler) handleSetAgentForDelayedOrder(w http.ResponseWriter, r *http.Request) {

	var (
		req  SetAgentForDelayedOrderRequest
		resp domain.AgentSetDelayedOrderResponse
		msg  string = "Done"
		err  error
	)

	if validateJsonRequest(w, r, &req) {

		var (
			domainReq                                    domain.AgentSetDelayedOrderRequest
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

		resp, err = a.service.SetAgentToDelayedOrder(r.Context(), domainReq)

		if err != nil {
			msg = err.Error()
		}
	}
	response.NewSuccess(msg, resp, 200).Send(w)
}
