package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Story struct {
    Title string `json: "title"`
    Url   string `json: "url"`
}

type TopStoriesResponse []int

func getStory(c chan Story, id int) {
    url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
    
    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body)
    
    var story Story
    
    _ = json.Unmarshal(body, &story)
    
    c <- story
}

func getTopStoryIds() TopStoriesResponse {
    resp, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
    body, _ := ioutil.ReadAll(resp.Body)
    
    var topStoriesResponse TopStoriesResponse
    
    _ = json.Unmarshal(body, &topStoriesResponse)
    
    topStoriesResponse = topStoriesResponse[:50]
    
    return topStoriesResponse
}

func main() {
    topStoryIds := getTopStoryIds()
    
    fmt.Printf("Got %d hits\n", len(topStoryIds))
    
    c := make(chan Story)
    
    for _, id := range topStoryIds {
        go getStory(c, id)
    }
    
    for story := range c {
        fmt.Printf("%s\n", story.Title)
    }
}
