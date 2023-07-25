package main

import (
	"fmt"
	"github.com/omid-h70/shop-service/cmd"
	"github.com/omid-h70/shop-service/internal/adapter/handler"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/core/service"
	"path/filepath"
	"runtime"
	"testing"
)

func ProjectRootDir() string {
	_, d, _, _ := runtime.Caller(1)
	return filepath.Dir(d)
}

var (
	dbTestClientConfig_1 = repository.MySqlConfig{
		DbServerAddr: "localhost",
		DbServerPort: "3306",
		DbName:       "shop_service_db",
		DbUser:       "root",
		DbPass:       "omid2142", //work workbench pass
		//DbPass:       "secret",//homepass
	}
	testServerConfig_1 = cmd.ServerConfig{
		Addr: "0.0.0.0",
		Port: "8000",
	}
)

func Test_main(t *testing.T) {

	//Setting Database
	serverConfig := testServerConfig_1
	dbClientConfig := dbTestClientConfig_1

	fmt.Println("Db Config", dbClientConfig)
	fmt.Println("App Config", dbClientConfig)

	appDbClient := repository.NewRepositoryMySqlDB(dbClientConfig)

	appAgentService := service.NewAgentService(repository.NewAgentRepositoryMySqlDB(appDbClient), 1000)
	appOrderService := service.NewOrderService(repository.NewOrderRepositoryMySqlDB(appDbClient), 1000)
	appVendorService := service.NewVendorService(repository.NewVendorRepositoryMySqlDB(appDbClient), 1000)

	cmd.NewAppConfig().
		ServerAddress(serverConfig).
		RegisterService(appOrderService, appAgentService, appVendorService).
		Run(handler.NewGorillaMuxRouter())

	fmt.Println("Hi, i'm up")
}
