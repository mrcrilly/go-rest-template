package teehee

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
	RequestCount    int64       `json:"request_count"`
	HttpStatusCodes map[int]int `json:"http_status_codes"`
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
