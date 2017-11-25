package teehee

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var globalStatus *Status
var globalLogger *logrus.Logger

func Init(configFrom string) (err error) {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)

	// set some sensible defaults so we can ignore
	// the need for a configuration file
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.ip", "127.0.0.1")
	viper.SetDefault("logging.enabled", true)
	viper.SetDefault("logging.file", "app.log")

	// attempt to load configuration
	viper.SetConfigName(configFrom)
	viper.AddConfigPath(".")

	// ignoring the error because we've set sensible(?)
	// defaults
	viper.ReadInConfig()

	if viper.GetBool("logging.enabled") {
		globalLogger = logrus.New()
		fd, err := os.OpenFile(viper.GetString("logging.file"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		globalLogger.Out = fd
	}

	return
}
