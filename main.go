package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var globalConfig *Configuration
var globalStatus *Status
var globalLogger *logrus.Logger

func init() {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)
	globalConfig = new(Configuration)
	globalLogger = logrus.New()
}

func main() {
	err := globalConfig.Load("config.json")
	checkErrorAndPanic(err)

	fd, err := os.OpenFile(globalConfig.Logging.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func serveRequests() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/config", HandlerReadOnlyConfig)
	router.HandleFunc("/health", HandlerHealthCheck)

	server := new(http.Server)

	if globalConfig.Http.Tls {
	} else {
		server.Addr = fmt.Sprintf("%s:%s",
			globalConfig.Http.BindIp, globalConfig.Http.BindPort,
		)
		server.Handler = router
		err = server.ListenAndServe()
	}

	return
}
