package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var globalStatus *Status
var globalLogger *logrus.Logger

func init() {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)
	globalLogger = logrus.New()

	viper.SetDefault("http_bind_port", "8080")
	viper.SetDefault("http_bind_addr", "127.0.0.1")
	viper.SetDefault("logging_file", "app.log")
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	checkErrorAndPanic(err)

	fd, err := os.OpenFile(viper.GetString("logging_file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErrorAndPanic(err)
	globalLogger.Out = fd

	err = serveRequests()
	checkErrorAndPanic(err)
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}
