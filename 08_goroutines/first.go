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

func (c *Computer) StartTimer(t time.Duration) {
    fmt.Println("Starting timer...")
    time.Sleep(t)
    fmt.Println("Time up!")
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    t := 3 * time.Second
    
    // TODO: Spawn this function in a goroutine
    computer.StartTimer(t)
    
    // IGNORE THIS:
    // This is a hack, so the program doesn't quit before the goroutine finishes
    time.Sleep(10 * time.Second)
}