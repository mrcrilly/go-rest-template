package teehee

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
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

func HandlerReadOnlyConfig(w http.ResponseWriter, r *http.Request) {
	go globalStatus.IncrementRequestCount()
	go globalLogger.WithFields(logrus.Fields{
		"handler": "healthcheck",
		"method":  r.Method,
		"path":    r.URL.String(),
	}).Info("working...")

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	result.Result = viper.AllSettings()

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
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
