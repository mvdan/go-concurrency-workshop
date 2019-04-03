package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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

type result struct {
	url     string
	elapsed time.Duration
	status  string
	size    int
}

func main() {
	ch := make(chan result)

	// fire one goroutine per url
	for _, url := range urls {
		go get(ch, url)
	}

	// get one result back for each of the urls
	var results []result
	for range urls {
		// receive a result from whichever of the spawned goroutines
		// that's finished; blocks until any one finishes.
		res := <-ch
		results = append(results, res)
	}

	// now, sort the results list from fastest to slowest
	sort.Slice(results, func(i, j int) bool {
		return results[i].elapsed < results[j].elapsed
	})

	// finally, print the results
	for _, res := range results {
		fmt.Println(res.url, res.elapsed, res.status, res.size)
	}
}

func get(ch chan result, url string) {
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
	// send the result back to the main goroutine via the channel
	ch <- result{url, elapsed, resp.Status, len(body)}
}
