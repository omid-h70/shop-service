package handler

import (
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
)

type VendorHandler struct {
	service service.VendorService
}

type VendorReportRequest struct {
}

func (a *VendorHandler) handleGetAllDelayReportsByVendor(w http.ResponseWriter, r *http.Request) {

	var req domain.VendorReportRequest
	resp, _ := a.service.GetAllDelayedOrdersByVendor(r.Context(), req)

	outData := "Done"
	response.NewSuccess(outData, resp, 200).Send(w)
}

func NewVendorHandler(service service.VendorService) VendorHandler {
	return VendorHandler{service: service}
}
