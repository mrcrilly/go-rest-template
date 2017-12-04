package teehee

import (
	"errors"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

var globalStatus = NewStatus()
var globalLogger *logrus.Logger

func SetLogger(l *logrus.Logger) {
	if l == nil {
		panic(errors.New("need a valid logger object"))
	}
	globalLogger = l
}

// GetRouter will construct an httprouter
// configured with our endpoints and ready to
// be used for serving traffic
func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", HandlerIndex)
	router.GET("/config", HandlerReadOnlyConfig)
	router.GET("/health", HandlerHealthCheck)
	return router
}
