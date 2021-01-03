package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Not Accessible:", url)
		c <- url
		return
	}

	fmt.Println("Accessible:", url)
	c <- url
}

func checkLinks(links []string, ongoing bool) {
	c := make(chan string)
	for _, ln := range links {
		go checkLink(ln, c)
	}

	if ongoing {
		for {
			go func(url string) {
				time.Sleep(time.Second)
				checkLink(url, c)
			}(<-c)
		}
		/*Alternate syntax:
		for l := range c {
			go func(url string) {
				time.Sleep(time.Second)
				checkLink(url, c)
			}(l)
		}
		*/
	} else {
		for idx := 0; idx < len(links); idx++ {
			<-c
		}
	}
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	checkLinks(links, false)
	fmt.Println("Ran once. Switching to ongoing...")
	checkLinks(links, true)
}
