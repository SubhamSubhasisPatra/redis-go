package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "might be down"
		return
	}
	fmt.Println(link, "is active")
	c <- "it up !"
}

func main() {
	fmt.Println("Go Channels  ... ")
	links := []string{
		"http://amazon.com",
		"http://google.com",
		"http://flipkart.com",
		"http://walmart.com",
		"http://facebook.com",
	}
	// create a cahnnel
	c := make(chan string)

	for _, val := range links {
		go checkLink(val, c)
		// fmt.Println(<-c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}
