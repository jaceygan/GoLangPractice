package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://github.com",
		"http://facebook.com",
		"http://porkbun.com",
		"http://storygraph.com",
	}

	fmt.Println("Checking the following urls:")

	for _, u := range urls {
		fmt.Println(u)
	}
	fmt.Println("Here we go.....")
	fmt.Println()
	for _, u := range urls {
		printStatus(u)
	}
}

func printStatus(u string) {
	_, err := http.Get(u)

	if err != nil {
		fmt.Println(u, "is down.. :(")
	}
	fmt.Println(u, "is up! :)")

}
