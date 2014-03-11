/*
  Para los usuarios de Java la paqueteria no indica  la ruta fisica que se acostumbra a crear en los archivos .java, en GOLANG.
  la paqueteria es un nombre l-gico que no necesariamente debe estar en creada en el disco duro.
*/
package 1basicconcepts

// El paquete fmt que puede ser encontrado en http://golang.org/pkg/fmt/ para su an-lisis  es un paquete que permite formatear la escritura y entrada de cadenas en una consola.
// El paquete net que puede ser encontrado en http://golang.org/pkg/net/ y es un paquete para el manejo de mensajes de entrada y salida en la red.
// El paquete os que puede ser encontrada en http://golang.org/pkg/os/ y es un paquete que permite escribir y leer archivos en el disco duro independiente de la plataforma(Linux, Unix, Mac, Windows) que se este ejecutando GOLANG
// El paquete time puede ser encontrado en http://golang.org/pkg/time/ y es un paquete que permite realizar mediones y  calculos de tiempo 
// Todos los paquetes de GOLANG pueden ser encontrados en http://golang.org/pkg/, la anterior ruta es el equivalente a la API de JAVA
import (
    "fmt"
    "net"
    "os"
    "time"
    "file"
)

// En GOLANG no existen clases como en otros lenguajes(Java), en GOLANG solo existen funciones y estructuras, para este caso solo se define una funci-n, las estructuras se ver-an en otros ejemplos.
// recordar que la funci-n con nombre main es especial y es el punto de entrada a una aplicaci-n, el equivalente en Java a el m-todo"public static void main(String[] arg)"
func exercice() {
    
    // debido a que ya importamos lso paquetes fmt,net,os y time para usarlos solo se debe volver colocar el nombre del paquete seguido de un punto y la primera letra de la funci-n a invocar debe estar en mayuscula
    
    fmt.Println("Bienvenido a ejercicio 2 del Tour de GOLANG!")

    fmt.Println("La fecha es : ", time.Now())

    fmt.Println("Si se intenta abrir un archivo que no existe, GOLANG nos imprimir- que no encontro ningun archivo")
    fmt.Println(os.Open("filename"))
    
    fmt.Println("Si se intenta abrir un archivo que existe, GOLANG imprimir- un valor que indicara la referencia del objeto :")
    fmt.Println(os.Open("2file.txt"))
    
    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
	    //log.Fatal(err)
    }
    fmt.Printf("read %d bytes: %q\n", count, data[:count])

    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "google.com"))
}