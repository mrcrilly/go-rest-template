package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var globalConfig *Configuration
var globalStatus *Status

func init() {
	globalStatus = new(Status)
	globalConfig = new(Configuration)

	err := globalConfig.Load("config.json")
	if err != nil {
		panic(err)
	}

}

func main() {
	err := serveRequests()
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
		if err != nil {
			return
		}
	}

	return
}
