package memory

import (
    "time"
    "strconv"
    
    )
    


func InsertArrayInt(param int) ([]Usuario, time.Duration) {

    //var arrayUsuarios [param]Usuario
    
    arrayUsuarios := make([]Usuario, param)

    tiempo1 := time.Now()
  //  mString = make(map[string]Usuario)

    for i := 0; i < param; i++ {
        str := "name" 
        str += strconv.Itoa(i)
        str2 := "dir" 
        str2 += strconv.Itoa(i)
        
        arrayUsuarios[i] = Usuario{ str, str2}
    }
    return   arrayUsuarios , time.Now().Sub(tiempo1)
}