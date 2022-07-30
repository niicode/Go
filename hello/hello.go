package main

import (
    "fmt"
    "log"

    "example.com/greeting"
)

func main() {
    log.SetPrefix("greeting: ")
    log.SetFlags(0)

    message, err := greeting.Hello("")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}