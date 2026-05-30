package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://archlinux.org",
		"http://golang.org",
		"http://google.com",
		"http://reddit.com",
		"http://wikipedia.org",
		"http://youtube.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkWebLink(link, c)
	}

	for l := range c {
		go func(lin string) {
			time.Sleep(2 * time.Second)
			checkWebLink(lin, c)
		}(l)
	}
}

func checkWebLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down!")
		c <- l
		return
	}

	fmt.Println(l, "is up and running!")
	c <- l
}
