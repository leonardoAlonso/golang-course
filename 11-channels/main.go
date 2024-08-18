package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"https://www.google.com",
		"https://facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.udemy.com",
		"https://www.coursera.org",
		"https://www.linkedin.com",
	}

	c := make(chan string) // create new channel

	for _, url := range websites {
		go checkLinkl(url, c) // pass channel to function
	}

	for l := range c {
		go checkLinkl(l, c) // receive message from channel
	}
}

func checkLinkl(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // send message to channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // send message to channel
}
