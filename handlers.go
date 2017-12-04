package teehee

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// HandlerIndex is the default route and the
// entry point into the service.
func HandlerIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go globalStatus.IncrementRequestCount()
	go globalLogger.WithFields(logrus.Fields{
		"handler": "healthcheck",
		"method":  r.Method,
		"path":    r.URL.String(),
	}).Info("working...")

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}

// HandlerHealthCheck allows a monitoring suites to
// detect if the service is active or not, or if
// it's handled any HTTP status codes that respresent
// problems.
func HandlerHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go globalStatus.IncrementRequestCount()
	go globalLogger.WithFields(logrus.Fields{
		"handler": "healthcheck",
		"method":  r.Method,
		"path":    r.URL.String(),
	}).Info("working...")

	result := new(Message)
	result.Message = "alive"
	result.Status = http.StatusOK
	result.Result = globalStatus

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}
