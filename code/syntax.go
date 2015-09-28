package main

import "fmt"


func add(a int, b int) int {
  return a + b
}

func main() {
  a := 1
  b := 3

  c := add(a, b)

  // Print the result
  fmt.Printf("%d + %d = %d", a, b, c)
}
