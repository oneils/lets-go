package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create a newTestApplication helper which returns an instance of our
// application struct containing mocked dependencies.
func newTestApplication(t *testing.T) *application {
	// create app with logger, because logRequest and recoverPanic use loggers.
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
}

type testServer struct {
	*httptest.Server
}

// newTestServer is a helper function to testing server creation
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

// get makes a GET request to a given url path using the test server client, and returns the
// response status code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
