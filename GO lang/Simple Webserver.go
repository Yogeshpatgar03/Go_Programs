package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    var port string
    fmt.Print("Enter the port to listen on (e.g., 8080): ")
    fmt.Scan(&port)

    server := &http.Server{
        Addr:    ":" + port,
        Handler: mux,
    }

    log.Printf("Server listening on port %s...\n", port)
    log.Fatal(server.ListenAndServe())
}
