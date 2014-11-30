package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	env := r.URL.Path[1:]
	fmt.Fprintf(w, "Hello there, my time is %v\n\nValue of environment variable %v is %v", time.Now(), env, os.Getenv(env))
}

func main() {
	http.HandleFunc("/", pingHandler)
	panic(http.ListenAndServe(":8080", nil))
}
