package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func buildGetStoryUrl(id int) string {
    return fmt.Sprintf(
        "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty",
        id,
    )
}

func main() {
    url := buildGetStoryUrl(8863)
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    
    fmt.Println(string(body))
}