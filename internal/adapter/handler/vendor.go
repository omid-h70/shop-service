package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/omid-h70/shop-service/internal/core/service"
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrInvalidApplicationType   = errors.New("Request Application Type Must be json")
	ErrInvalidCardNumberPattern = errors.New("Card Number Is invalid")
	ErrInvalidDataBaseQuery     = errors.New("DataBase Error")
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

	var req VendorReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" &&
		r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
		response.NewError(ErrInvalidApplicationType, http.StatusBadRequest).Send(w)
		return
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	//Fix Arabic Persian
	//fixArabicPersianStrings(&req)
	//
	////Check Card Number Pattern
	//if !helper.CheckCardNumber(req.CardFromNum) || !helper.CheckCardNumber(req.CardToNum) {
	//	response.NewError(ErrInvalidCardNumberPattern, http.StatusBadRequest).Send(w)
	//	return
	//}
	//
	////Mapping Request To Domain
	//inAmount, parseErr := strconv.ParseInt(req.TransactionAmount, 10, 64)
	//if parseErr != nil {
	//	response.NewError(err, http.StatusBadRequest).Send(w)
	//	return
	//}
	//
	//input := domain.NewTransaction()
	//input.SetCardFromInfo(domain.NewCard(req.CardFromNum))
	//input.SetCardToInfo(domain.NewCard(req.CardToNum))
	//input.SetAmount(inAmount)
	//
	//accountListInfo, err := a.service.ExecuteCardTransfer(r.Context(), input)
	//if err != nil || len(accountListInfo) != 2 {
	//	response.NewError(err, http.StatusBadRequest).Send(w)
	//	return
	//}

	/*
		cardFromInfoOut := accountListInfo[0]
		cardToInfoOut := accountListInfo[1]
		var sender string
		senderReceptorList := []string{cardFromInfoOut.CustomerInfo.PhoneNum}
		receiverReceptorList := []string{cardToInfoOut.CustomerInfo.PhoneNum}

		senderMsg := a.notifyService.GetSenderNotifyMessage(cardFromInfoOut, WITHDRAW_FILE_TEMPLATE)
		receiverMsg := a.notifyService.GetReceiverNotifyMessage(cardToInfoOut, DEPOSIT_FILE_TEMPLATE)

		go func() {
			a.notifyService.SendNotifyMessage(sender, senderReceptorList, senderMsg)
			a.notifyService.SendNotifyMessage(sender, receiverReceptorList, receiverMsg)
		}()

		outData := map[string]map[string]string{
			"SenderMsg": {
				"To":  cardFromInfoOut.CustomerInfo.PhoneNum,
				"Msg": senderMsg,
			},
			"ReceiverMsg": {
				"To":  cardToInfoOut.CustomerInfo.PhoneNum,
				"Msg": receiverMsg,
			},
			"status": {
				"Message": "Done",
			},
		}
	*/
	outData := "Done"
	response.NewSuccess(outData, "", 200).Send(w)
}

//
//func NewAccountHandler(service service.TransactionService, notifyService domain.PushNotificationService) AccountHandler {
//	return AccountHandler{service: service}
//}
