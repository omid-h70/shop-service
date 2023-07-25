package handler

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/core/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	ADD_NEW_ORDER_URL = "/add_new_order"
	MOCK_URL
)

func Test_delay_report_should_fail_when_order_id_or_vendor_id_is_invalid(t *testing.T) {
	var jsonData = []byte(`{
		"vendor_id" : "123456",
		"order_id": "123456",
	}`)

	router := mux.NewRouter()

	app := NewOrderHandler(service.NewOrderService(repository.NewOrderRepositoryMockDB(), 1000), MOCK_URL)
	router.HandleFunc(ADD_NEW_ORDER_URL, app.handleMakeDelayReport).Methods(http.MethodPost)

	request, _ := http.NewRequest(http.MethodPost, ADD_NEW_ORDER_URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Error("Test Failed")
	}
}

func Test_should_generate_eta_from_mock_service(t *testing.T) {

}
