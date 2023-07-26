package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MysqlContainerName string `mapstructure:"MYSQL_CONTAINER_NAME"`
	MysqlServerAddr    string `mapstructure:"MYSQL_SERVER_ADDR"`
	MysqlDBName        string `mapstructure:"MYSQL_DATABASE"`
	MysqlUser          string `mapstructure:"MYSQL_USER"`
	MysqlPass          string `mapstructure:"MYSQL_PASS"`
	MysqlRootPass      string `mapstructure:"MYSQL_ROOT_PASS"`
	MysqlHostPort      string `mapstructure:"MYSQL_HOST_PORT"`
	MysqlContainerPort string `mapstructure:"MYSQL_CONTAINER_PORT"`

	AppContainerName string `mapstructure:"APP_CONTAINER_NAME"`
	AppName          string `mapstructure:"APP_NAME"`
	AppServerAddr    string `mapstructure:"APP_SERVER_ADDR"`
	AppHostPort      string `mapstructure:"APP_HOST_PORT"`
	AppContainerPort string `mapstructure:"APP_CONTAINER_PORT"`
}

func LoadConfig(path string, cfg *Config) (err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env") // json, xml , whatever

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	fmt.Println(cfg)
	return
}
