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
    fmt.Fprintln(res, "Benchmark Innova4j")
    fmt.Fprintln(res, "")


    fmt.Fprintln(res, "<h2>Prueba  de carga con mapas de String</h2>")
    fmt.Fprintln(res, "100000 structuras")
    m , a := memory.InsertString(100000)
    fmt.Fprintln(res, a)
    
    b := memory.AccessString(m);
    m = nil
    fmt.Fprintln(res, b)
    
    fmt.Fprintln(res, "500000 structuras")
    m1 , a1 := memory.InsertString(500000)
    fmt.Fprintln(res, a1)
    b1 := memory.AccessString(m1);
    m1 = nil
    fmt.Fprintln(res, b1)
    
    fmt.Fprintln(res, "1000000 structuras")
    m2 , a2 := memory.InsertString(1000000)
    fmt.Fprintln(res, a2)
    b2 := memory.AccessString(m2);
    m2 = nil
    fmt.Fprintln(res, b2)
    
    fmt.Fprintln(res, "5000000 structuras")
    m3 , a3 := memory.InsertString(50000000)
    fmt.Fprintln(res, a3)
    b3 := memory.AccessString(m3);
    m3 = nil
    fmt.Fprintln(res, b3)
    
    fmt.Fprintln(res, "<h2>Prueba  de carga con mapas de Int</h2>")
    fmt.Fprintln(res, "100000 structuras")
    m4 , a4 := memory.InsertInt(100000)
    fmt.Fprintln(res, a4)
    b4 := memory.AccessInt(m4);
    m4 = nil
    fmt.Fprintln(res, b4)
    
    fmt.Fprintln(res, "500000 structuras")
    m5 , a5 := memory.InsertInt(500000)
    fmt.Fprintln(res, a5)
    b5 := memory.AccessInt(m5);
    m5 = nil
    fmt.Fprintln(res, b5)
    
    fmt.Fprintln(res, "1000000 structuras")
    m6 , a6 := memory.InsertInt(1000000)
    fmt.Fprintln(res, a6)
    b6 := memory.AccessInt(m6);
    m6 = nil
    fmt.Fprintln(res, b6)
    
    fmt.Fprintln(res, "5000000 structuras")
    m7 , a7 := memory.InsertInt(50000000)
    fmt.Fprintln(res, a7)
    b7 := memory.AccessInt(m7);
    m7 = nil
    fmt.Fprintln(res, b7)
   
    fmt.Fprintln(res, "fin test")
}