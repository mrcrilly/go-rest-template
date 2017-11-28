package teehee

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func init() {
	ConfigFromString("")
}

type ExpectedResult struct {
	Verb        string
	Path        string
	Body        io.Reader
	Response    *Message
	ResultIsNil bool
}

func TestHandlerIndex(t *testing.T) {
	expected := &ExpectedResult{
		"GET",
		"/",
		nil,
		&Message{
			Message: "ok",
			Status:  200,
			Result:  nil,
		},
		true,
	}

	req, err := http.NewRequest(expected.Verb, expected.Path, expected.Body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := httprouter.New()
	handler.GET(expected.Path, HandlerIndex)
	handler.ServeHTTP(rr, req)

	result := MessageJsonToObject(rr.Body)
	if result.Status != expected.Response.Status {
		t.Errorf("Status code in response (%d) doesn't match what we expected (%d)",
			result.Status, expected.Response.Status)
	}

	if result.Message != expected.Response.Message {
		t.Errorf("Message in response doesn't match what was expected: %s\n", result.Message)
	}

	if result.Result != nil && expected.ResultIsNil {
		t.Error("Result should be nil for this endpoint")
	}
}
