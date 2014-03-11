/*
  Para los usuarios de Java la paqueteria no indica  la ruta fisica que se acostumbra a crear en los archivos .java, en GOLANG.
  la paqueteria es un nombre lógico que no necesariamente debe estar en creada en el disco duro.
*/
package 1basicconcepts

// El paquete fmt que puede ser encontrado en http://golang.org/pkg/fmt/ para su an-lisis  es un paquete que permite formatear la escritura y entrada de cadenas en una consola.
// Todos los paquetes de GOLANG pueden ser encontrados en http://golang.org/pkg/, la anterior ruta es el equivalente a la API de JAVA
import "fmt"

// En GOLANG no existen clases como en otros lenguajes(Java), en GOLANG solo existen funciones y estructuras, para este caso solo se define una funci-n, las estructuras se ver-an en otros ejemplos.
// recordar que la funci-n con nombre main es especial y es el punto de entrada a una aplicaci-n, el equivalente en Java a el m-todo"public static void main(String[] arg)"
func exercice() {
    // debido a que ya importamos el paquete para usarlo se debe volver colocar el nombre del paquete seguido de un punto y la primera letra de la funci-n a invocar debe estar en mayuscula
    fmt.Println("Hello my firts line on GOLANG")
}