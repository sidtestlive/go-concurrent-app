package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Simulate some work
    time.Sleep(1 * time.Second)
    fmt.Fprintf(w, "Handled by pod at %s\n", time.Now().Format(time.RFC3339))
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
