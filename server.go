package teehee

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func StartServer() (err error) {
	router := httprouter.New()
	router.GET("/", HandlerIndex)
	router.GET("/config", HandlerReadOnlyConfig)
	router.GET("/health", HandlerHealthCheck)

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
