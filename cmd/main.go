package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mrcrilly/teehee"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.ip", "127.0.0.1")
	viper.SetDefault("logging.enabled", true)
	viper.SetDefault("logging.file", "app.log")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/teehee/")
	viper.AddConfigPath("$HOME/.teehee")
	viper.AddConfigPath(".")
	checkErrorAndPanic(viper.ReadInConfig())

	if viper.GetBool("logging.enabled") {
		fd, err := os.OpenFile(
			viper.GetString("logging.file"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		checkErrorAndPanic(err)
		l := logrus.New()
		l.Out = fd
		teehee.SetLogger(l)
	}

	return
}

func startServer() (err error) {
	router := teehee.GetRouter()
	httpServer := new(http.Server)
	if viper.GetBool("http.tls.enabled") {
	} else {
		httpServer.Addr = fmt.Sprintf("%s:%s",
			viper.GetString("http.ip"),
			viper.GetString("http.port"),
		)
		httpServer.Handler = router
		err = httpServer.ListenAndServe()
	}

	return
}

func main() {
	startServer()
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}
