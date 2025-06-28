package dhru

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
)

func FindApi(inputURL string) (string, error) {
	// Ensure URL has https:// scheme
	inputURL = ensureHttpsScheme(inputURL)

	// Parse the URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Define variants to try
	subdomains := []string{"", "www.", "api."}
	endpoints := []string{"/", "/api/index.php", "/connect/api/index.php"}

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use channels for results and completion signaling
	resultChan := make(chan string, 1)
	done := make(chan struct{})

	// Start parallel checks
	var wg sync.WaitGroup
	for _, subdomain := range subdomains {
		for _, endpoint := range endpoints {
			wg.Add(1)
			go tryEndpoint(ctx, &wg, resultChan, parsedURL, subdomain, endpoint)
		}
	}

	// Signal when all goroutines complete
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for success or completion
	select {
	case serverUrl := <-resultChan:
		return serverUrl, nil
	case <-done:
		return "", fmt.Errorf("no https API endpoint found for %s - tried %d variations", parsedURL.Host, len(subdomains)*len(endpoints))
	}
}

// Helper function to ensure URL has https:// scheme
func ensureHttpsScheme(inputURL string) string {
	if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
		return "https://" + inputURL
	} else if strings.HasPrefix(inputURL, "http://") {
		return "https://" + inputURL[7:]
	}
	return inputURL
}

// Helper function to try a specific endpoint
func tryEndpoint(ctx context.Context, wg *sync.WaitGroup, resultChan chan string, parsedURL *url.URL, subdomain, endpoint string) {
	defer wg.Done()

	// Check if context is canceled
	select {
	case <-ctx.Done():
		return
	default:
		// Continue with the check
	}

	testURL := fmt.Sprintf("%s://%s%s%s", parsedURL.Scheme, subdomain, parsedURL.Host, endpoint)

	realServerUrl, err := url.Parse(testURL)
	if err != nil {
		return // Skip if URL parsing fails
	}

	username := "example"
	apikey := "keyExample"

	repository := newRepo(realServerUrl, username, apikey)
	if repository.ping() {
		// Send result to channel and ignore if channel is full
		select {
		case resultChan <- testURL:
		default:
		}
	}
}
