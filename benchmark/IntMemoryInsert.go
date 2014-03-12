package memory

import (
    "time"
    )


func InsertInt(param int) (map[int64]Usuario, time.Duration) {

    tiempo1 := time.Now()
    m = make(map[int64]Usuario)


    for i := 0; i < param; i++ {
        str := "name" 
        str2 := "dir" 
        m[int64(i)] = Usuario{ str, str2}
    }
    return   m , time.Now().Sub(tiempo1)
}