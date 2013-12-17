package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// simplified to the function from thread here to use a buffered channel to limit requests: https://groups.google.com/forum/#!topic/golang-nuts/r-Ye2v5BB0A
func Crawl(url string, depth int, fetcher Fetcher) {
	reqch := make(chan int)
	workch := make(chan bool, 4) // limit the number of concurrent requests with a buffered channe;
	visited := make(map[string]bool)

	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		fmt.Println("Crawling", url)
		defer func() { reqch <- -1 }()
		if _, ok := visited[url]; ok {
			return
		}
		if depth <= 0 {
			return
		}
		visited[url] = true
		workch <- true
		defer func() { <-workch }()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		reqch <- len(urls)
		for _, u := range urls {
			go crawl(u, depth-1)
		}
	}
	go crawl(url, depth)

	actsum := 1
	for diff := range reqch {
		actsum += diff
		if actsum == 0 {
			break
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
