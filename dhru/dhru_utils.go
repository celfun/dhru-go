package dhru

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
)

func FindApi(inputURL string) (string, error) {
	// Check if URL already has a scheme
	if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
		// Add https:// prefix if no scheme exists
		inputURL = "https://" + inputURL
	} else if strings.HasPrefix(inputURL, "http://") {
		// Replace http:// with https://
		inputURL = "https://" + inputURL[7:]
	}

	// Parse the URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	subdomains := []string{"", "www.", "api."}
	endpoints := []string{"/", "/api/index.php", "/connect/api/index.php"}

	type result struct {
		serverUrl string
		err       error
	}

	// Create a channel for results
	resultChan := make(chan result, 1)
	// Create a channel to signal we're done checking all endpoints
	done := make(chan struct{})
	// WaitGroup to track all goroutines
	var wg sync.WaitGroup

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Try all combinations in parallel
	for _, subdomain := range subdomains {
		for _, endpoint := range endpoints {
			wg.Add(1)
			go func(sd, ep string) {
				defer wg.Done()

				// Check if context is canceled
				select {
				case <-ctx.Done():
					return
				default:
					// Continue with the check
				}
				testURL := fmt.Sprintf("%s://%s%s%s", parsedURL.Scheme, sd, parsedURL.Host, ep)
				realServerUrl, err := url.Parse(testURL)
				if err != nil {
					return // Skip if URL parsing fails
				}
				username := "example"
				apikey := "keyExample"

				repository := newRepo(realServerUrl, username, apikey)
				if repository.ping() {
					// Send result to channel and cancel other goroutines
					select {
					case resultChan <- result{testURL, nil}:
						cancel() // Cancel all other goroutines
					default:
						// Result already received, do nothing
					}
				}
			}(subdomain, endpoint)
		}
	}

	// Close done channel when all goroutines complete
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for either a successful result or all goroutines to finish
	select {
	case r := <-resultChan:
		return r.serverUrl, r.err
	case <-done:
		return "", fmt.Errorf("could not find a valid endpoint")
	}
}
