package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var urls = []string{
	"https://github.com/",
	"https://govern.cat/",
	"https://hackupc.com/",
	"https://www.fib.upc.edu/",
	"https://golang.org/",
	"https://google.com/",

	// uncomment to do objective #3
	// "https://github.com/does-not-exist",
}

func main() {
	for _, url := range urls {
		start := time.Now()

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		elapsed := time.Since(start)
		fmt.Println(url, elapsed, resp.Status, len(body))
	}
}
