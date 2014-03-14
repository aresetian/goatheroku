package concurrency

import "fmt"

func sum(a []int, c chan int) {
    sum := <- c
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func suma2( a , b int, c chan int) {
    z := <-c
}

func main() {
    
    a := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], <-c)
    
   
    
    x, y := 17 , <-c // receive from c
    fmt.Println(x, y, x+y)
}