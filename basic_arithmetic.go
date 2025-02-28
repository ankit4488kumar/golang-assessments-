// basic_arithmetic.go - Perform basic arithmetic operations on two numbers
package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)  //We take two integers as input using fmt.Scan()

    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    if b != 0 {
        fmt.Println("Division:", a/b)
    } else {
        fmt.Println("Division by zero is not allowed")
    }
}
