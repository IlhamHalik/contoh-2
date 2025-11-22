package main

import "fmt"

func main() {
    var n int
    
    fmt.Print("input n: ")
    fmt.Scan(&n)

    hasil := 0
    for i := 1; i <= n; i+=1 {
        hasil += i
    }

    fmt.Println(hasil)
}
