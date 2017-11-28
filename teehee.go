package teehee

import (
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var globalStatus *Status
var globalLogger *logrus.Logger

func config(configFrom io.Reader) (err error) {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)

	// set some sensible defaults so we can ignore
	// the need for a configuration file
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.ip", "127.0.0.1")
	viper.SetDefault("logging.enabled", true)
	viper.SetDefault("logging.file", "app.log")
	err = viper.ReadConfig(configFrom)
	if err != nil {
		return err
	}

	if viper.GetBool("logging.enabled") {
		fd, err := os.OpenFile(
			viper.GetString("logging.file"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		globalLogger = logrus.New()
		if err != nil {
			return err
		}
		globalLogger.Out = fd
	}

	return
}

func ConfigFromString(configFrom string) (err error) {
	return config(strings.NewReader(configFrom))
}

func ConfigFromReader(configFrom io.Reader) (err error) {
	return config(configFrom)
}

func Config(configFrom string) (err error) {
	fd, err := os.Open(configFrom)
	if err != nil {
		return
	}
	return config(fd)
}
