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

func (c *Computer) StartTimer(channel chan string, t time.Duration) {
    fmt.Println("Starting timer...")
    time.Sleep(t)
    channel <- "Time up!"
}

func (c *Computer) StartInterval(channel chan string, t time.Duration) {
    fmt.Println("Starting interval...")
    
    for i := 0; i < 10; i++ {
        time.Sleep(t)
        channel <- fmt.Sprintf("Interval %d up!", i)
    }

    close(channel)
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    channel := make(chan string)
    
    fmt.Printf("How long? ")
    
    // TODO: Get the number of seconds from user input
    
    t := time.Duration(seconds) * time.Second
    go computer.StartInterval(channel, t)
    
    for msg := range channel {
        fmt.Println(msg)
    }
    
    fmt.Println("Exited")
}