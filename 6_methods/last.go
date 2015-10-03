package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }

    computer.Describe()
}