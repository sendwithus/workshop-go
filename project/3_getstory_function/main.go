package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func getStory(id int) string {
    url := buildGetStoryUrl(id)
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    
    return string(body)
}

func buildGetStoryUrl(id int) string {
    return fmt.Sprintf(
        "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty",
        id,
    )
}

func main() {
    fmt.Println(getStory(8863))
}