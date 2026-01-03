package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		panic(err)
	}

	fmt.Println("ID:", post.ID)
	fmt.Println("Title", post.Title)
}
