package main

import (
	"bytes"
	"github.com/oneils/lets-go/internal/assert"
	"io"
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
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
