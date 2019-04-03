# Go concurrency workshop

First, download a recent version of Go. Either through your package manager, or
by downloading 1.12.1 from https://golang.org/doc/install.

### TL;DR introduction to Go

It's a statically typed language that's very similar to C++ and Java. It has:

* Simple C-like syntax
* Basic data structures: int, bool, string, list (slice), map
* Garbage collection, similar to Java
* Concurrency support via goroutines (lightweight threads) and channels

We'll jump right into the exercise. If you encounter issues, use google and
stackoverflow, or just ask people near you.

### Starting up

This repository contains the skeleton for a concurrency exercise in Go.

Ensure you're on Go 1.12 or later:

	$ go version
	go version go1.12.1 linux/amd64

Build it and run it:

	$ go run main.go

You can use `go build` to build a binary. You can use `goimports` to
automatically format the code and add missing imports:
https://godoc.org/golang.org/x/tools/cmd/goimports

### Initial version

The first version sequentially fetches a number of HTTP URLs, printing each of
them, how long it took to fetch them, the status code, and the size of the
returned body.

We'll make multiple changes to this file. Edit the `main.go` file. You can make
commits locally if you want, to keep track of your progress.

### Objective #1: fetch all URLs concurrently

Doing this work sequentially is wasteful. We can fetch all the URLs at once. Use
goroutines to do that: https://tour.golang.org/concurrency/1

With a goroutine, you can run a piece of code asynchronously, without needing to
wait for it to finish.

Gotcha: the program will finish as soon as the initial (main) goroutine
finishes. What's the best way to fix that? You can start with a `time.Sleep`: https://golang.org/pkg/time/

### Objective #2: sort URL results by body size

If you implemented the concurrent program correctly, you should see the URLs
that took less time to fetch first.

Modify the program to instead order the results, so that the first is the one
with the smallest body size.

You'll need to use channels, where each of the spawned goroutines sends back its
result to the main goroutine. Example: https://tour.golang.org/concurrency/2

### Objective #3: handle errors gracefully

We might encounter an error when fetching some of the URLs. Uncomment the last
one, which should return a 404 not found.

The current program panics immediately on any error. Instead, send the error as
part of the result via the channel to the main goroutine, and print all the
error results at the very end of the output.

### Addendum: Gotchas

* No semicolons: they aren't necessary
* Types are reversed: `func(urls []string)`, not `func(string[] urls)`
* Unlike C++/Java, funcs can return many values: `func() (url string, ok
  bool)`
* Builds and imports are based on packages, similar to Java
* Unused imports are variables are build errors
