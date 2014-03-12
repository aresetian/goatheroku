package memory

import (
    "time"
    "fmt"
    "strconv"
    
    )
type Usuario struct {
    name, dir string
}

var m map[string]Usuario

func PruebaMemoria(param int) (map[string]Usuario, time.Duration) {

    tiempo1 := time.Now()
    m = make(map[string]Usuario)
   /* {
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }*/


    for i := 0; i < param; i++ {
        str := "name" 
        str += strconv.Itoa(i)
        str2 := "dir" 
        str2 += strconv.Itoa(i)
        key := "key" 
        key += strconv.Itoa(i)
        m[key] = Usuario{ str, str2}
    }
    
    // debido a que ya importamos el paquete para usarlo se debe volver colocar el nombre del paquete seguido de un punto y la primera letra de la funci-n a invocar debe estar en mayuscula
    /*
    for i := 0; i < param; i++ {
        key := "key" 
        key += string(i)
        
        fmt.Println("pruebaMemoria500mil : " , m[key])
    }*/
    
    //tiempo2 := time.Second
    
    fmt.Printf("pruebaMemoria %d tiempo total : " , param , time.Now().Sub(tiempo1))
    
    //respuesta := tiempo2 - tiempo1
    return   m , time.Now().Sub(tiempo1)
}