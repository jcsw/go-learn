package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type blogPost struct {
	UserID  int    `json:"userId"`
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"body"`
}

func main() {

	var httpTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var httpClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: httpTransport,
	}

	httpGET(httpClient)

	httpPOST(httpClient)
}

func httpGET(httpClient *http.Client) {

	response, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("Could not read response\n%v\n", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read response body\n%v\n", err)
		return
	}

	blogPosts := []blogPost{}
	err = json.Unmarshal(body, &blogPosts)
	if err != nil {
		fmt.Printf("Could not parse 'blogPosts'\n%v\n", err)
		return
	}

	for i, blogPost := range blogPosts {
		fmt.Printf("BlogPost [%d] [%+v]\n", i, blogPost)
	}

}

func httpPOST(httpClient *http.Client) {

	blogPost := blogPost{UserID: 1, Title: "Golang http", Content: "Hello world!"}
	postBody, err := json.Marshal(blogPost)
	if err != nil {
		fmt.Printf("Could not convert blogPost to json\n%v\n", err)
	}

	response, err := httpClient.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("Could not read response\n%v\n", err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("Status code returnuded: %v\n", response.StatusCode)
}
