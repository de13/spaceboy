package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hostname:") })

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	osHostname, _ := os.Hostname()
	osHostname = "Hostname: " + osHostname + "\n"
	getHostname := string(writer.Body.Bytes()[:])
	if osHostname != getHostname {
		t.Errorf("Hostname should be %v but got %q", osHostname, writer.Body.Bytes())
	}
}

func TestReady(t *testing.T) {
	t.Parallel()
	ready = 30
	readiness := check{"ready.html", ready}
	mux := http.NewServeMux()
	mux.HandleFunc("/ready", ready.state)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ready", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 503 {
		t.Errorf("Response code is %v, should be 503", writer.Code)
	}
	ready = 0
	time.Sleep(1 * time.Second)
	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/ready", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v, should be 200", writer.Code)
	}
}

func TestHealthz(t *testing.T) {
	t.Parallel()
	healthz = 30
	healthiness := check{"strangerThings.html", healthz}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health.state)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthz", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v, should be 200", writer.Code)
	}
	healthz = 0
	time.Sleep(1 * time.Second)
	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/healthz", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 503 {
		t.Errorf("Response code is %v, should be 503", writer.Code)
	}
}
