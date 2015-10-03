package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }

    fmt.Println(computer)
}