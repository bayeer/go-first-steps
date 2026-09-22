package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	baseURL := "https://api.frankfurter.dev/v2/rates"
	// 1. parse baseURL.
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	// 2. add query params.
	q := u.Query()
	q.Set("date", time.Now().Format("2006-01-02"))
	q.Set("base", "USD")
	q.Set("quotes", "RUB")
	u.RawQuery = q.Encode()

	// 3. get full URL
	url := u.String()

	// 4. create cancel context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5. create request with context.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 6. set headers.
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	// 7. create http client.
	client := &http.Client{Timeout: time.Second * 10}

	// 8. do the request.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 9. read response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
