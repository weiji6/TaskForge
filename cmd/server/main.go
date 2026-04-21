package main

import "log"
import "net/http"

func main() {
    log.Println("Starting TaskForge server...")
    http.ListenAndServe(":8080", nil)
}