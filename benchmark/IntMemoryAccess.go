package memory

import (
    "time"
    //"fmt"
    
    )
type Usuario struct {
    name, dir string
}

var m map[int64]Usuario

func AccessInt(mapa map[int64]Usuario)  time.Duration {

    tiempo1 := time.Now()
   
 
    for i := 0; i < len(mapa); i++ {
        
        str :=  mapa[int64(i)]
        mapa[int64(i)] = str
        //fmt.Println("AccessInt : " , mapa[int64(i)])
    }
    
    return   time.Now().Sub(tiempo1)
}