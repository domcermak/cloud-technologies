package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// get()
	// post()
	// customClient()
	// timeout()
}

func get() {
	resp, err := http.Get("http://localhost:8080/dump")
	if err != nil {
		panic(err)
	}

	// !!!
	defer resp.Body.Close()

	_ = resp.Header
	_ = resp.Status
	_ = resp.ContentLength
	_ = resp.StatusCode
}

func post() {
	resp, err := http.Post("http://localhost:8080/post", "text/plain", strings.NewReader("my name"))
	if err != nil {
		panic(err)
	}

	// !!!
	defer resp.Body.Close()
}

func customClient() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,
		MaxIdleConns:        150,
		MaxIdleConnsPerHost: 150,
		IdleConnTimeout:     time.Minute,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	httpClient := &http.Client{Timeout: 15 * time.Second, Transport: transport}

	_, _ = httpClient.Head("http://localhost:8080/dump")
}

func timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/sleeper", http.NoBody)
	if err != nil {
		panic(err)
	}

	_, err = http.DefaultClient.Do(req)
	fmt.Println("done", err)
}
