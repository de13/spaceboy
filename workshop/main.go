package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pod, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Hostname:", pod)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	if uptime() > time.Second*120 {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	t, _ := template.ParseFiles("strangerThings.html")
	t.Execute(w, uptime() < time.Second*120)
}

func ready(w http.ResponseWriter, r *http.Request) {
	if uptime() < time.Second*30 {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	t, _ := template.ParseFiles("ready.html")
	t.Execute(w, uptime() > time.Second*30)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/ready", ready)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
