package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/aresetian/goatheroku/benchmark"
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
   // fmt.Fprintln(res, "hello, world equipo innova otro")
    fmt.Fprintln(res, "Benchmark Innova4j")
//    fmt.Fprintln(res, basicconcepts.Exercice())
//    fmt.Fprintln(res, basicconcepts.Exercice2())
    fmt.Fprintln(res, "Prueba con 100000 structuras")
    fmt.Fprintln(res, benchmark.PruebaMemoria100mil())
    fmt.Fprintln(res, "fin test")
}