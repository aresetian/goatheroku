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
    fmt.Fprintln(res, "Prueba  de carga con 100000 structuras")
    //fmt.Fprintln(res, memory.PruebaMemoria100mil())
    m , a := memory.PruebaMemoria(100000)
    m= nil
    fmt.Fprintln(res, m)
    fmt.Fprintln(res, a)
    
    fmt.Fprintln(res, "Prueba  de carga con 500000 structuras")
    m1 , a1 := memory.PruebaMemoria(100000)
    m1= nil
    fmt.Fprintln(res, m1)
    fmt.Fprintln(res, a1)
    
    fmt.Fprintln(res, "Prueba  de carga con 1000000 structuras")
    m2 , a2 := memory.PruebaMemoria(1000000)
    m2= nil
    fmt.Fprintln(res, m2)
    fmt.Fprintln(res, a2)
    //fmt.Fprintln(res, memory.PruebaMemoria500mil(100000))
    //fmt.Fprintln(res, memory.PruebaMemoriaMillon())
    //fmt.Fprintln(res, memory.PruebaMemoria500mil(2))
    fmt.Fprintln(res, "fin test")
}