package main

import (
	"fmt"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	go globalStatus.IncrementRequestCount()
	go globalLogger.Println("index received request")

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerReadOnlyConfig(w http.ResponseWriter, r *http.Request) {
	go globalStatus.IncrementRequestCount()

	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	result.Result = globalConfig

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	go globalStatus.IncrementRequestCount()

	result := new(Message)
	result.Message = "alive"
	result.Status = http.StatusOK
	result.Result = globalStatus

	go globalStatus.IncrementHttpStatusCode(result.Status)
	fmt.Fprint(w, result.ToJsonString())
}
