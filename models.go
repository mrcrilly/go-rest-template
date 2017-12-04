package teehee

import (
	"bytes"
	"encoding/json"
	"io"
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

func MessageJsonToObject(m io.Reader) *Message {
	jDecoder := json.NewDecoder(m)
	var result Message
	_ = jDecoder.Decode(&result)
	return &result
}

type Status struct {
	Lock sync.Mutex `json:"-"`

	// request statistics
	RequestCount    int64       `json:"request_count"`
	HttpStatusCodes map[int]int `json:"http_status_codes"`
}

// NewStatus is primarily for initialising the
// map more than anything. Just a concenience method
func NewStatus() *Status {
	return &Status{
		HttpStatusCodes: make(map[int]int, 0),
	}
}

// IncrementHttpStatusCode will update the status map
// in a thread safe manner so we can track the number
// of, and type of, HTTP status codes we're returning
// to clients
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
