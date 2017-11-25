package teehee

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func StartServer() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/config", HandlerReadOnlyConfig)
	router.HandleFunc("/health", HandlerHealthCheck)

	server := new(http.Server)

	if viper.GetBool("http.tls.enabled") {
	} else {
		server.Addr = fmt.Sprintf("%s:%s",
			viper.GetString("http.ip"), viper.GetString("http.port"),
		)
		server.Handler = router
		err = server.ListenAndServe()
	}

	return
}
