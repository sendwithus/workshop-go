package main

import "fmt"

func isLongerThan(message string, n int) bool {
    if len(message) > n {
        return true
    } else {
        return false
    }
}

func main() {
    fmt.Println(isLongerThan("Hello World", 30))
    fmt.Println(isLongerThan("Hello World", 3))
}