package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// START CODE OMIT
type elapsedTimeClient struct {
	client *http.Client
}

func (etc *elapsedTimeClient) Get(url string) (*http.Response, error) {
	start := time.Now()
	res, err := etc.client.Get(url)
	res.Header.Set("X-Elapsed-Time", fmt.Sprintf("%d", time.Since(start)))
	return res, err
}

func doTimedGet(url string) {
	cl := &elapsedTimeClient{&http.Client{}} // HL
	res, err := cl.Get(url)
	if err != nil {
		log.Printf("Request was an error: %s", err)
		return
	}

	log.Println(res.StatusCode)
	log.Println(res.Header.Get("X-Elapsed-Time"))
}
// END CODE OMIT

// START INTERFACE OMIT
type Getter interface {
	Get() (*http.Response, error) // HL
}

// http.ClientでもelapsedTimeClientでも渡せる！
func doGet(g Getter, url string) {
	res, err := g.Get(url)
	if err != nil {
		log.Printf("Request was an error: %s", err)
		return
	}
	log.Println(res.StatusCode)
}
// END INTERFACE OMIT

