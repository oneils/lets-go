package main

import (
	"bytes"
	"github.com/oneils/lets-go/internal/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// example of unit testing of handlers
func TestPingUnit(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}

// example of end-to-end test
func TestPing(t *testing.T) {
	// create app with logger, becase logRequest and recoverPanic use loggers.
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	// The network address that the test server is listening on is contained in
	// the ts.URL field. We can  use this along with the ts.Client().Get() method
	// to make a GET /ping request against the test server. This returns a
	// http.Response struct containing the response.
	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
