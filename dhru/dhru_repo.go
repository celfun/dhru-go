package dhru

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type repo struct {
	serverUrl url.URL
	username  string
	apiKey    string
}

func newRepo(serverUrl *url.URL, username string, apiKey string) *repo {
	return &repo{
		serverUrl: *serverUrl,
		username:  username,
		apiKey:    apiKey,
	}
}
func (r *repo) doRequest(action string) ([]byte, error) {
	// Create URL without query parameters
	u := r.serverUrl // Create a copy of the URL

	// Create values for parameters
	params := url.Values{}
	params.Set("requestformat", "JSON")
	params.Set("action", action)
	params.Set("username", r.username)
	params.Set("apiaccesskey", r.apiKey)

	// Create a new request with POST method, passing params in the body
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(params.Encode()))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Set the content type header for form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	if req.URL.String() == "https://www.directcodes.org/connect/api/index.php" {
		println("Request URL:", req.URL.String())
	}

	// Send the request using default client
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	body, err := io.ReadAll(resp.Body)

	var apiData apiResponse
	if err := json.Unmarshal(body, &apiData); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}
	if apiData.ERROR != nil {
		return nil, fmt.Errorf("error from server: %s", apiData.ERROR)
	}

	return body, nil
}

func (r *repo) ping() bool {
	_, err := r.doRequest("accountinfo")
	if err != nil {
		if strings.Contains(err.Error(), "error from server") {
			return true
		}
	}
	return false
}

func (r *repo) getAccountInfo() (accountDetails, error) {
	jsonData, err := r.doRequest("accountinfo")
	if err != nil {
		return accountDetails{}, err
	}
	var accountData apiResponse
	err = json.Unmarshal(jsonData, &accountData)
	if err != nil {
		return accountDetails{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	if accountData.ERROR != nil {
		return accountDetails{}, fmt.Errorf("api error: %s", accountData.ERROR[0].Message)
	}
	return accountData.SUCCESS[0].AccountInfo, nil
}

func (r *repo) getImeiServiceList() (map[string]group, error) {
	jsonData, err := r.doRequest("imeiservicelist")
	if err != nil {
		return nil, err
	}
	var root apiResponse
	err = json.Unmarshal(jsonData, &root)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return root.SUCCESS[0].List.Groups, nil
}
