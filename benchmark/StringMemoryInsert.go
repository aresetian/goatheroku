package memory

import (
    "time"
    "strconv"
    
    )
    
    var mString map[string]Usuario

func InsertString(param int) (map[string]Usuario, time.Duration) {

    tiempo1 := time.Now()
    mString = make(map[string]Usuario)

    for i := 0; i < param; i++ {
        str := "name" 
        str += strconv.Itoa(i)
        str2 := "dir" 
        str2 += strconv.Itoa(i)
        key := "key" 
        key += strconv.Itoa(i)
        mString[key] = Usuario{ str, str2}
    }
    return   mString , time.Now().Sub(tiempo1)
}