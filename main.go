package main

import (
	"fmt"
	"github.com/omid-h70/shop-service/cmd"
	"github.com/omid-h70/shop-service/internal/adapter/handler"
	"github.com/omid-h70/shop-service/internal/adapter/repository"
	"github.com/omid-h70/shop-service/internal/core/service"
	"os"
)

var (
	mockUrl = "https://run.mocky.io/v3/122c2796-5df4-461c-ab75-87c1192b17f7"
	//
	dbTestClientConfig = repository.MySqlConfig{
		DbServerAddr: "localhost",
		DbServerPort: "3306",
		DbName:       "shop_service_db",
		DbUser:       "root",
		DbPass:       "secret", //home pass
		//DbPass:       "omid2142", //work workbench pass
	}
	testServerConfig = cmd.ServerConfig{
		Addr: "0.0.0.0",
		Port: "8000",
	}
)

func main() {

	build := os.Getenv("BUILD_TYPE")
	//Setting Database
	dbClientConfig := repository.MySqlConfig{
		DbServerAddr: os.Getenv("MYSQL_CONTAINER_NAME"),
		DbServerPort: os.Getenv("MYSQL_CONTAINER_PORT"),
		DbName:       os.Getenv("MYSQL_DATABASE"),
		DbUser:       os.Getenv("MYSQL_USER"),
		DbPass:       os.Getenv("MYSQL_PASS"),
	}
	serverConfig := cmd.ServerConfig{
		Addr: os.Getenv("APP_SERVER_ADDR"),
		Port: os.Getenv("APP_HOST_PORT"),
	}

	if build == "local" {
		serverConfig = testServerConfig
		dbClientConfig = dbTestClientConfig
	}

	fmt.Println("App Starting With")
	fmt.Println("Build Config", build)
	fmt.Println("Server Config", serverConfig)
	fmt.Println("DB Config", serverConfig)

	appDbClient := repository.NewRepositoryMySqlDB(dbClientConfig)

	//Setting Services And Handlers
	appAgentService := service.NewAgentService(repository.NewAgentRepositoryMySqlDB(appDbClient), 1000)
	appOrderService := service.NewOrderService(repository.NewOrderRepositoryMySqlDB(appDbClient), 1000)
	appVendorService := service.NewVendorService(repository.NewVendorRepositoryMySqlDB(appDbClient), 1000)

	cmd.NewAppConfig().
		ServerAddress(serverConfig).
		RegisterService(appOrderService, appAgentService, appVendorService, mockUrl).
		Run(handler.NewGorillaMuxRouter())

}
