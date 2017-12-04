package teehee

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func init() {
	logger := logrus.New()
	SetLogger(logger)
}

var HandlerBasicsTable = []struct {
	Handler     httprouter.Handle
	Verb        string
	Path        string
	Body        io.Reader
	Response    *Message
	ResultIsNil bool
}{
	{
		Handler: HandlerHealthCheck,
		Verb:    "GET",
		Path:    "/config",
		Body:    nil,
		Response: &Message{
			Message: "alive",
			Status:  200,
		},
		ResultIsNil: false,
	},
	{
		Handler: HandlerReadOnlyConfig,
		Verb:    "GET",
		Path:    "/config",
		Body:    nil,
		Response: &Message{
			Message: "ok",
			Status:  200,
		},
		ResultIsNil: false,
	},
	{
		Handler: HandlerIndex,
		Verb:    "GET",
		Path:    "/",
		Body:    nil,
		Response: &Message{
			Message: "ok",
			Status:  200,
		},
		ResultIsNil: true,
	},
}

func TestHandlerBasics(t *testing.T) {
	for _, tt := range HandlerBasicsTable {
		req, err := http.NewRequest(tt.Verb, tt.Path, tt.Body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := httprouter.New()
		handler.GET(tt.Path, tt.Handler)
		handler.ServeHTTP(rr, req)

		result := MessageJsonToObject(rr.Body)
		if result.Status != tt.Response.Status {
			t.Errorf("Status code in response (%d) doesn't match what we expected (%d)",
				result.Status, tt.Response.Status)
		}

		if result.Message != tt.Response.Message {
			t.Errorf("Message in response doesn't match what was expected: %s\n", result.Message)
		}

		if result.Result != nil && tt.ResultIsNil {
			t.Error("Result should be nil for this endpoint")
		}
	}
}
