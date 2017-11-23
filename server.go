package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func serveRequests() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/config", HandlerReadOnlyConfig)
	router.HandleFunc("/health", HandlerHealthCheck)

	server := new(http.Server)

	if viper.GetBool("http_tls_enabled") {
	} else {
		server.Addr = fmt.Sprintf("%s:%s",
			viper.GetString("http_bind_ip"), viper.GetString("http_bind_port"),
		)
		server.Handler = router
		err = server.ListenAndServe()
	}

	return
}
