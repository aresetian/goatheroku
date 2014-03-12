package memory

import (
    "time"
    //"fmt"
    )

func AccessArrayInt(arreglo []Usuario)  time.Duration {

    tiempo1 := time.Now()
   
 
    for i := 0; i < len(arreglo); i++ {
        
        jota := arreglo[i] 
        arreglo[i] = jota
        //fmt.Println("AccessString : " , arreglo[i])
    }
    return   time.Now().Sub(tiempo1)
}