package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{
			url,
			err,
			0,
		}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{
			url,
			nil,
			t,
		}
		resp.Body.Close()
	}
}

func main() {
	ch := make(chan result)
	urls := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	for _, url := range urls {
		go get(url, ch)
	}

	for range urls {
		r := <-ch

		if r.err != nil {
			log.Printf("%s %s\n", r.url, r.err)
		} else {
			log.Printf("%s %s\n", r.url, r.latency)
		}
	}
}
