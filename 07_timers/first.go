package main

import (
    "fmt"
    "time"
)

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

// TODO: Make a StartTimer function that...
//  - Sleeps for a given duration
//  - Prints a message when done

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    t := 3 * time.Second
    computer.StartTimer(t)
}