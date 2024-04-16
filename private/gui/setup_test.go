package gui

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testMainWindow MainWindow

func TestMain(m *testing.M) {
	a := test.NewApp()
	testMainWindow = *NewMainWindow(a)
	testMainWindow.httpClient = client
	os.Exit(m.Run())
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})

var jsonToReturn = `{
  "ts": 1712972463502,
  "tsj": 1712972455479,
  "date": "Apr 12th 2024, 09:40:55 pm NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 2344.53,
      "xagPrice": 27.943,
      "chgXau": -32.47,
      "chgXag": -0.5345,
      "pcXau": -1.366,
      "pcXag": -1.8769,
      "xauClose": 2377,
      "xagClose": 28.4775
    }
  ]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
