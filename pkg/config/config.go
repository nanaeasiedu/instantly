package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Settings Conf

type Conf struct {
	Env               string
	UnityClientID     string
	UnityClientSecret string
	BrokerToken       string
	BrokerSender      string
	BrokerCallbackURL string
	BrokerBaseURL     string
	DBName            string
	DBPath            string
	MigrationsDir     string
	RedisURL          string
	RedisPassword     string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
	}

	Settings.Env = viper.GetString("SERVER_ENV")
	Settings.UnityClientID = viper.GetString("UNITY_CLIENT_ID")
	Settings.UnityClientSecret = viper.GetString("UNITY_CLIENT_SECRET")
	Settings.BrokerToken = viper.GetString("BROKER_TOKEN")
	Settings.BrokerSender = viper.GetString("BROKER_SENDER")
	Settings.BrokerBaseURL = viper.GetString("BROKER_BASE_URL")
	Settings.MigrationsDir = viper.GetString("MIGRATIONS_DIR")
	Settings.DBName = viper.GetString("DB_NAME")
	Settings.DBPath = viper.GetString("DB_PATH")
	Settings.BrokerCallbackURL = viper.GetString("BROKER_CALLBACK_URL")
	Settings.RedisURL = viper.GetString("REDIS_URL")
	Settings.RedisPassword = viper.GetString("REDIS_PASSWORD")
}
