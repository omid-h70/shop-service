package handler

import (
	"fmt"
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/domain"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
)

type VendorHandler struct {
	service service.VendorService
	//notifyService domain.PushNotificationService
}

type VendorReportRequest struct {
	//	CardFromNum       string `json:"card_from_number" validate:"required,len=16"`
	//	CardToNum         string `json:"card_to_number" validate:"required,len=16"`
	//	TransactionAmount string `json:"transaction_amount" validate:"required"`
}

func (a *VendorHandler) handleGetAllDelayReportsByVendor(w http.ResponseWriter, r *http.Request) {

	var req domain.VendorReportRequest
	resp, _ := a.service.GetAllDelayedOrdersByVendor(r.Context(), req)
	fmt.Println(resp)

	outData := "Done"
	response.NewSuccess(outData, "", 200).Send(w)
}

//
//func NewVendorHandler(service service.VendorService) VendorHandler {
//	return VendorHandler{service: service}
//}
