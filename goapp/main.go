package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var startTime time.Time
var healthz, ready int

func init() {
	startTime = time.Now()
	flag.IntVar(&healthz, "healthz", 120, "Delay (in seconds) during which the application is considered as healthy.")
	flag.IntVar(&healthz, "h", 120, "Delay (in seconds) during which the application is considered as healthy.")
	flag.IntVar(&ready, "ready", 30, "Initial delay (in seconds) during which application is considered as not ready.")
	flag.IntVar(&ready, "r", 30, "Initial delay (in seconds) during which application is considered as not ready.")
}

func main() {
	flag.Parse()
	pod, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	readiness := check{"ready.html", ready}
	healthiness := check{"strangerThings.html", healthz}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hostname:", pod) })
	http.HandleFunc("/healthz", healthiness.state)
	http.HandleFunc("/ready", readiness.state)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func uptime() time.Duration {
	return time.Since(startTime)
}

func (c check) state(w http.ResponseWriter, r *http.Request) {
	if uptime() > time.Second*time.Duration(healthz) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	t, _ := template.ParseFiles(filepath.Join("templates", c.template))
	t.Execute(w, uptime() < time.Second*time.Duration(c.delay))
}

type check struct {
	template string
	delay    int
}
