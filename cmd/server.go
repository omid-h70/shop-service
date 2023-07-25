package cmd

import (
	"fmt"
	"github.com/omid-h70/shop-service/internal/adapter/handler"
	"github.com/omid-h70/shop-service/internal/core/service"
)

type ServerConfig struct {
	Addr string
	Port string
}

type AppConfig struct {
	serverCnf  ServerConfig
	appHandler handler.AppHandler
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		appHandler: handler.NewAppHandler(),
	}
}

func (cnf *AppConfig) RegisterService(order service.OrderService,
	agent service.AgentService,
	vendor service.VendorService,
	mockurl string) *AppConfig {

	cnf.appHandler.RegisterAgentService(agent)
	cnf.appHandler.RegisterOrderService(order, mockurl)
	cnf.appHandler.RegisterVendorService(vendor)
	return cnf
}

//func (cnf *AppConfig) CustomerRepo(repo domain.CustomerRepository) *AppConfig {
//	cnf.repo = repo
//	return cnf
//}
//
//func (cnf *AppConfig) NotifyService(notifyRepo domain.PushNotificationService) *AppConfig {
//	cnf.notifyRepo = notifyRepo
//	return cnf
//}

func (cnf *AppConfig) ServerAddress(servConfig ServerConfig) *AppConfig {
	cnf.serverCnf = servConfig
	return cnf
}

func (cnf *AppConfig) Run(router handler.HandlerInterface) {
	router.SetAppHandlers(&cnf.appHandler)
	router.Listen(fmt.Sprintf("%s:%s", cnf.serverCnf.Addr, cnf.serverCnf.Port))
}
