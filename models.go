package teehee

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

// Message represents the response given to clients
// after making a request (successful or not) to the
// service's endpoints.
type Message struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// ToJsonString is a helper function for converting
// a Message from its native type into a string for
// easy printing or other types of rendering
func (m *Message) ToJsonString() string {
	buf := new(bytes.Buffer)
	jEncoder := json.NewEncoder(buf)
	_ = jEncoder.Encode(m)
	return buf.String()
}

// MessageJsonToObject is a helper function for
// converting Message types received as JSON into
// their native type.
func MessageJsonToObject(m io.Reader) *Message {
	jDecoder := json.NewDecoder(m)
	var result Message
	_ = jDecoder.Decode(&result)
	return &result
}

// Status is used for tracking various aspects
// of the service's state, such as requests handled
// and HTTP status codes returned.
//
// It's safe to update this struct on a service-by-service
// basis.
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

// IncrementRequestCount is a simple, thread safe
// method for incrementing the total number of requests
// this service has received
func (s *Status) IncrementRequestCount() {
	s.Lock.Lock()
	s.RequestCount += 1
	s.Lock.Unlock()
}
