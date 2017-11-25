package teehee

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var globalStatus *Status
var globalLogger *logrus.Logger

func Init(configFrom, logTo string) (err error) {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)
	globalLogger = logrus.New()

	// set some sensible defaults so we can ignore
	// the need for a configuration file
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.ip", "127.0.0.1")
	viper.SetDefault("logging.enabled", true)
	viper.SetDefault("logging.file", "app.log")

	// attempt to load configuration
	viper.SetConfigName(configFrom)
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	fd, err := os.OpenFile(viper.GetString("logging.file"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	globalLogger.Out = fd
	return
}
