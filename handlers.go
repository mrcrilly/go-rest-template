package main

import (
	"fmt"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	globalStatus.IncrementRequestCount()

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerReadOnlyConfig(w http.ResponseWriter, r *http.Request) {
	globalStatus.IncrementRequestCount()

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	result.Result = globalConfig
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	globalStatus.IncrementRequestCount()

	result := new(Message)
	result.Message = "alive"
	result.Status = http.StatusOK
	result.Result = globalStatus
	fmt.Fprint(w, result.ToJsonString())
}
