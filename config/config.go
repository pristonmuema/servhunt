package config

import (
	"fmt"
	"github.com/spf13/viper"
	"servhunt/infra/utils"
)

var (
	logger = utils.GetRootLogger()
)

type Config struct {
	Database struct {
		ConnectionPool     int    `json:"ConnectionPool"`
		User               string `json:"User"`
		Password           string `json:"Password"`
		Name               string `json:"Name"`
		ConnectionUrl      string `json:"ConnectionUrl"`
		MaxIDleConnections int    `json:"MaxIdleConnections"`
	} `json:"Database"`
	Cache struct {
		User          string `json:"User"`
		Database      int    `json:"Database"`
		ConnectionUrl string `json:"ConnectionUrl"`
		Password      string `json:"Password"`
	} `json:"Cache"`
}

func InitViperConfig() (config *Config) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("config")
	v.AddConfigPath("/config")
	v.AddConfigPath("./config")
	v.AddConfigPath("../config")

	//Override with env
	_ = v.BindEnv("Database.User", "DB_USER")
	_ = v.BindEnv("Database.Password", "DB_PASSWORD")
	_ = v.BindEnv("Database.ConnectionUrl", "DB_CONNECTION_URL")
	_ = v.BindEnv("Database.Name", "DB_NAME")
	//load cache configs
	_ = v.BindEnv("Cache.Password", "REDIS_PASSWORD")
	_ = v.BindEnv("Cache.Url", "REDIS_URL")

	err := v.ReadInConfig()
	err = v.Unmarshal(&config)

	if err != nil {
		logger.Panic(fmt.Sprintf("No configuration file loaded - using defaults %s", err.Error()))
	}
	return
}
