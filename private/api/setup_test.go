package api

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

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
