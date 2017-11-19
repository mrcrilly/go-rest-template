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
	Lock sync.Mutex `json:"-"`

	// request statistics
	RequestCount int64 `json:"request_count"`

	HttpStatusCodes map[int]int `json:"http_status_codes"`
	// status codes
	//Http200s int64 `json:"http_200s"` // ok
	//Http301s int64 `json:"http_301s"` // perm redirect
	//Http302s int64 `json:"http_302s"` // redirect; or should this be 303/307 now?
	//Http400s int64 `json:"http_400s"` // bad request
	//Http401s int64 `json:"http_401s"` // unauthorized
	//Http403s int64 `json:"http_403s"` // forbidden
	//Http404s int64 `json:"http_404s"` // not found
	//Http405s int64 `json:"http_405s"` // method not allowed (get, post, etc)
	//Http415s int64 `json:"http_415s"` // unsupport media type (json, png, etc)
	//Http500s int64 `json:"http_500s"` // internal server error
	//Http502s int64 `json:"http_502s"` // bad gateway
	//Http503s int64 `json:"http_503s"` // service unavailable
}

func (s *Status) IncrementHttpStatusCode(code int) {
	s.Lock.Lock()
	if _, OK := s.HttpStatusCodes[code]; OK {
		s.HttpStatusCodes[code]++
	} else {
		s.HttpStatusCodes[code] = 1
	}
	s.Lock.Unlock()
}

func (s *Status) IncrementRequestCount() {
	s.Lock.Lock()
	s.RequestCount += 1
	s.Lock.Unlock()
}
