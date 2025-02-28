package main

import "fmt"

func fibonacci(n int) {
    a, b := 0, 1
    fmt.Print("Fibonacci Series: ")
    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}

func main() {
    var n int
    fmt.Print("Enter number of terms: ")
    fmt.Scan(&n)
    
    fibonacci(n)
}
