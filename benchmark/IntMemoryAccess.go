package memory

import (
    "time"
    "fmt"
    
    )
type Usuario struct {
    name, dir string
}

var m map[int]Usuario

func AccessInt(mapa map[int]Usuario)  time.Duration {

    tiempo1 := time.Now()
   
 
    for i := 0; i < len(mapa); i++ {
        fmt.Println("AccessInt : " , mapa[i])
    }
    
    return   time.Now().Sub(tiempo1)
}