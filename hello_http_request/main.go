package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true"
// @51.8297919,107.4510174,46430

func main() {
	baseURL := "https://api.open-meteo.com/v1/forecast"

	// 1. parse baseURL.
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	// 2. add query params.
	p := u.Query()
	p.Add("latitude", "51.827919")
	p.Add("longitude", "107.5610174")
	p.Add("current_weather", "true")
	u.RawQuery = p.Encode()

	url := u.String()
	// Print the full URL
	// println(url)
	// fmt.Println(url)

	// 3. create cancel context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 4. create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 5. set headers.
	req.Header.Set("Accept", "application/json")

	// 6. create http client.
	client := &http.Client{Timeout: 10 * time.Second}

	// 7. make request.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 8. read response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Status: %s", resp.Status)
	fmt.Println(string(body))
}
