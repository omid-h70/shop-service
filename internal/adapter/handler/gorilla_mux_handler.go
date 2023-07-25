package handler

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/omid-h70/shop-service/internal/adapter/response"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type muxRouter struct {
	router *mux.Router
}

func (mux *muxRouter) SetAppHandlers(appH *AppHandler) {

	api := mux.router.PathPrefix("/v1").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(defaultHandler)

	api.HandleFunc("/add_new_order", appH.HandleAddNewOrder).Methods(http.MethodPost)
	api.HandleFunc("/delay_report", appH.handleMakeDelayReport).Methods(http.MethodPost)
	api.HandleFunc("/get_all_delay_reports", appH.handleGetAllDelayReportsByVendor).Methods(http.MethodGet)
	api.HandleFunc("/set_agent", appH.handleSetAgentForDelayedOrder).Methods(http.MethodPost)
	api.HandleFunc("/handle_delayed_order", appH.handleDelayedOrderByAgent).Methods(http.MethodPost)
	api.HandleFunc("/health", healthCheck)
}

func (mux *muxRouter) Listen(serverAddr string) {

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         serverAddr,
		Handler:      mux.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed")
	}

	log.Fatal("Service down")
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	response.NewError(errors.New("Invalid request"), http.StatusBadRequest).Send(w)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	response.NewSuccess("Yo I'm up", "", http.StatusOK).Send(w)
}

func NewGorillaMuxRouter() *muxRouter {
	return &muxRouter{
		router: mux.NewRouter(),
	}
}
