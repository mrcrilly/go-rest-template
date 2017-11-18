package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

func (m *Message) ToJsonString() string {
	buf := new(bytes.Buffer)
	jEncoder := json.NewEncoder(buf)
	_ = jEncoder.Encode(m)
	return buf.String()
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	fmt.Fprint(w, result.ToJsonString())
}

func HandlerReadOnlyConfig(w http.ResponseWriter, r *http.Request) {
	result := new(Message)
	result.Message = "ok"
	result.Status = http.StatusOK
	result.Result = globalConfig
	fmt.Fprint(w, result.ToJsonString())
}
