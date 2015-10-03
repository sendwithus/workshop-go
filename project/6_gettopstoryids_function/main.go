package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Story struct {
    By    string `json: "by"`
    Title string `json: "title"`
    Url   string `json: "url"`
}

func (s *Story) Print() {
    fmt.Printf("%s: %s [%s]\n", s.By, s.Title, s.Url)
}

func getTopStoryIds() []int {
    resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    
    var ids []int
    _ = json.Unmarshal(body, &ids)
    
    return ids
}

func getStory(id int) Story {
    url := buildGetStoryUrl(id)
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    
    var story Story
    _ = json.Unmarshal(body, &story)
    
    return story
}

func buildGetStoryUrl(id int) string {
    return fmt.Sprintf(
        "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty",
        id,
    )
}

func main() {
    ids := getTopStoryIds()
    fmt.Printf("Got %d top stories\n", len(ids))
    
    topId := ids[0]
    
    story := getStory(topId)
    story.Print()
}