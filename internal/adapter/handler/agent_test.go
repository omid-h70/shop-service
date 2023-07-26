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
	ADD_SET_AGENT_URL = "/add_new_order"
)

func Test_delay_report_should_fail_when_agent_id_is_invalid(t *testing.T) {
	var jsonData = []byte(`{
		"agent_id" : "44444",
	}`)

	router := mux.NewRouter()

	app := NewAgentHandler(service.NewAgentService(repository.NewAgentRepositoryMockDB(), 1000))
	router.HandleFunc(ADD_NEW_ORDER_URL, app.handleSetAgentForDelayedOrder).Methods(http.MethodPost)

	request, _ := http.NewRequest(http.MethodPost, ADD_SET_AGENT_URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Error("Test Failed")
	}
}
