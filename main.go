package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/user-agent", userAgentHandler)
    http.HandleFunc("/poem", poemHandler)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello")
}

func userAgentHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User-Agent: %s\n", r.UserAgent())
    fmt.Fprintf(w, "Request Method: %s\n", r.Method)
}

func poemHandler(w http.ResponseWriter, r *http.Request) {
    poem := r.URL.Query().Get("poem")
    if poem == "" {
        http.Error(w, "Please provide a poem query parameter", http.StatusBadRequest)
        return
    }

    // Generate short poem
    // TODO: generate the rest of it

    fmt.Fprint(w, "Short poem generated")
}
