package main

import (
    "fmt"
    "log"
    "net/http"
)
            
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, It is test app!")
}
    
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8088", nil))
}
