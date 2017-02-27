package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Settings Conf

type Conf struct {
	Env               string
	APIClientUsername string
	APIClientPassword string
	UnityClientID     string
	UnityClientSecret string
	BrokerToken       string
	BrokerSender      string
	BrokerBaseURL     string
	PostgresURL       string
	PostgresDBName    string
	MigrationsDir     string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
	}

	Settings.Env = viper.GetString("SERVER_ENV")
	Settings.APIClientUsername = viper.GetString("API_CLIENT_USERNAME")
	Settings.APIClientPassword = viper.GetString("API_CLIENT_PASSWORD")
	Settings.UnityClientID = viper.GetString("UNITY_CLIENT_ID")
	Settings.UnityClientSecret = viper.GetString("UNITY_CLIENT_SECRET")
	Settings.BrokerToken = viper.GetString("BROKER_TOKEN")
	Settings.BrokerSender = viper.GetString("BROKER_SENDER")
	Settings.BrokerBaseURL = viper.GetString("BROKER_BASE_URL")
	Settings.PostgresURL = viper.GetString("POSTGRES_URL")
	Settings.PostgresDBName = viper.GetString("POSTGRES_DBNAME")
	Settings.MigrationsDir = viper.GetString("MIGRATIONS_DIR")
}
