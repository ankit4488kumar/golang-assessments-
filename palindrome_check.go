package main

import "fmt"

func isPalindrome(s string) bool {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scan(&str)

    if isPalindrome(str) {
        fmt.Println(str, "is a Palindrome")
    } else {
        fmt.Println(str, "is NOT a Palindrome")
    }
}