package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	pod, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Hostname:",  pod)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
        if err != nil {
          panic(err)
        }
}
