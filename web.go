package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/aresetian/goatheroku/tourgolang/basicconcepts"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world equipo innova otro")
    fmt.Fprintln(res, basicconcepts.Exercice())
    fmt.Fprintln(res, basicconcepts.Exercice2())
    fmt.Fprintln(res, "fin test")
}