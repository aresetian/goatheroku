package memory

import (
    "time"
    //"fmt"
    "strconv"
    
    )

func AccessString(mapa map[string]Usuario)  time.Duration {

    tiempo1 := time.Now()
   
 
    for i := 0; i < len(mapa); i++ {
        key := "key" 
        key += strconv.Itoa(i)
       // fmt.Println("AccessString : " , mapa[key])
    }
    return   time.Now().Sub(tiempo1)
}