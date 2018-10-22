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
	pod, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hostname:", pod) })

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
	readiness := check{"ready.html", 1, http.StatusServiceUnavailable, http.StatusOK}
	mux := http.NewServeMux()
	mux.HandleFunc("/ready", readiness.liveState)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ready", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 503 {
		t.Errorf("Response code is %v, should be 503", writer.Code)
	}
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
	healthiness := check{"strangerThings.html", 1, http.StatusOK, http.StatusServiceUnavailable}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthiness.healthState)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthz", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v, should be 200", writer.Code)
	}
	time.Sleep(1 * time.Second)
	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/healthz", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 503 {
		t.Errorf("Response code is %v, should be 503", writer.Code)
	}
}
