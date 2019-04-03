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
	"https://github.com/does-not-exist",
}

type result struct {
	url     string
	elapsed time.Duration
	status  string
	size    int
	err     error
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
		res1, res2 := results[i], results[j]
		// if both errors are nil or non-nil, sort by elapsed
		if (res1.err == nil) == (res2.err == nil) {
			return results[i].elapsed < results[j].elapsed
		}
		// otherwise, we want the result without an error first
		return res1.err == nil
	})

	// finally, print the results
	for _, res := range results {
		if res.err != nil {
			fmt.Println(res.url, res.elapsed, res.status, res.err)
		} else {
			fmt.Println(res.url, res.elapsed, res.status, res.size)
		}
	}
}

func get(ch chan result, url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- result{err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		// this case isn't covered by the error above
		ch <- result{
			url:     url,
			elapsed: time.Since(start),
			err:     fmt.Errorf("http status code %d", resp.StatusCode),
			status:  resp.Status,
		}
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- result{
			url:     url,
			elapsed: time.Since(start),
			err:     err,
			status:  resp.Status,
		}
		return
	}

	// send the result back to the main goroutine via the channel
	ch <- result{
		url:     url,
		elapsed: time.Since(start),
		status:  resp.Status,
		size:    len(body),
	}
}
