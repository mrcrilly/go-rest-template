package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var globalConfig *Configuration

func init() {
	globalConfig = new(Configuration)

	err := globalConfig.Load("config.json")
	if err != nil {
		panic(err)
	}

}

func main() {
	serveRequests()
}

func serveRequests() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/config", HandlerReadOnlyConfig)

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
