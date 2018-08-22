package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
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
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz/", funcHealthz)
	http.HandleFunc("/ready/", funcReady)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

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

func funcHealthz(w http.ResponseWriter, r *http.Request) {
	// Change the status code to 503 after x seconds
	if uptime() > time.Second*time.Duration(healthz) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	t, _ := template.ParseFiles("strangerThings.html")
	// Change the content of the web page to "crash" after x seconds
	t.Execute(w, uptime() < time.Second*time.Duration(healthz))
}

func funcReady(w http.ResponseWriter, r *http.Request) {
	// Set the status code to 503 during the first x seconds
	if uptime() < time.Second*time.Duration(ready) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	t, _ := template.ParseFiles("ready.html")
	// Set the content of the web page to "Ready" after x seconds
	t.Execute(w, uptime() > time.Second*time.Duration(ready))
}
