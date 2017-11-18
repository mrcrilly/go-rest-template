package main

import (
	"bytes"
	"encoding/json"
	"sync"
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

type Status struct {
	Lock         sync.Mutex `json:"-"`
	RequestCount int64      `json:"request_count"`
}

func (s *Status) IncrementRequestCount() {
	s.Lock.Lock()
	s.RequestCount += 1
	s.Lock.Unlock()
}
