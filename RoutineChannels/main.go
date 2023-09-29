package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://github.com",
		"http://facebook.com",
		"http://porkbun.com",
		"http://storygraph.com",
		"http://asasasfew.pop",
	}

	c := make(chan string)

	fmt.Println("Checking the following urls:")

	for _, u := range urls {
		fmt.Println(u)
	}
	fmt.Println("Here we go.....")
	fmt.Println()

	//initial run
	for _, u := range urls {
		go printStatus(u, c)
	}

	//subsequent runs based on response from channel

	for l := range c { //keep listening for response from channel c
		go func(link string) { //function literal (lambda for python)
			time.Sleep(time.Second * 5)
			printStatus(link, c)
		}(l)
	}

}

func printStatus(u string, c chan string) {
	_, err := http.Get(u)

	if err != nil {
		fmt.Println(u, "is down down down down!!!")
		fmt.Println(err)

	} else {
		fmt.Println(u, "is up and running")

	}
	c <- u
}
