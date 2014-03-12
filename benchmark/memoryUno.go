package memory


type Usuario struct {
    name, dir string
}
var m := map[string]Usuario

func PruebaMemoria100mil() string {

    tiempo1 := time.Second
    m = make(map[string]Usuario)
   /* {
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }*/


    for i := 0; i < 100000; i++ {
        m["" + i] = Usuario{"name" + i, "dir" + 1}
    }
    
    // debido a que ya importamos el paquete para usarlo se debe volver colocar el nombre del paquete seguido de un punto y la primera letra de la funci-n a invocar debe estar en mayuscula
    
    for i := 0; i < 100000; i++ {
        fmt.Println("pruebaMemoria100mil : "m["" + i])
    }
    
    tiempo2 := time.Second
    
    fmt.Println("pruebaMemoria100mil tiempo total : "  (tiempo2 - tiempo1))
    
    return   (tiempo2 - tiempo1)
}