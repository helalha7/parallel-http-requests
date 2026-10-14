package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)

	start := time.Now()
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
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
		"http://localhost:8080",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, url := range urls {
		go get(ctx, url, ch)
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
