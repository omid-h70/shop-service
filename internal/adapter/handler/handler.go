package handler

import (
	"encoding/json"
	"errors"
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
)

type HandlerInterface interface {
	SetAppHandlers(*AppHandler)
	Listen(Addr string)
}

type AppHandler struct {
	VendorHandler
	OrderHandler
	AgentHandler
}

var (
	ErrInvalidJsonFormat      = errors.New("Invalid Json Format")
	ErrInvalidJsonID          = errors.New("Invalid ID Format")
	ErrInvalidApplicationType = errors.New("Invalid Application Type")
)

func validateJsonRequest(w http.ResponseWriter, r *http.Request, data any) bool {

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return false
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" &&
		r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
		response.NewError(ErrInvalidApplicationType, http.StatusBadRequest).Send(w)
		return false
	}
	return true
}

func (appH *AppHandler) RegisterOrderService(service service.OrderService, mockurl string) {
	appH.OrderHandler.service = service
	appH.mockUrl = mockurl
}

func (appH *AppHandler) RegisterAgentService(service service.AgentService) {
	appH.AgentHandler.service = service
}

func (appH *AppHandler) RegisterVendorService(service service.VendorService) {
	appH.VendorHandler.service = service
}

func NewAppHandler() AppHandler {
	return AppHandler{}
}
