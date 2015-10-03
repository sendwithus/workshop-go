package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty")
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    
    fmt.Println(string(body))
}