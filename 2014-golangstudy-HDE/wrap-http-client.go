package main

import (
	"fmt"
	"net/http"
	"time"
)

// START GET OMIT
type myClient struct{ http.Client }

func (cl myClient) Get(url string) (*http.Response, error) {
	start := time.Now()
	resp, err := cl.Client.Get(url) // 無名でも型名でアクセスできる // HL
	if err != nil {
		return nil, err
	}
	resp.Header.Set("X-Elapsed-Time", fmt.Sprintf("%d", time.Since(start)))
	return resp, nil
}
// END GET OMIT
