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
func (r *repo) makeApiRequest(action string, dhruParams map[string]string) ([]byte, error) {
	var sb strings.Builder
	if dhruParams != nil {
		sb.WriteString("<PARAMETERS>")
		for key, value := range dhruParams {
			escapedValue := value
			sb.WriteString(fmt.Sprintf("<%s>%s</%s>", key, escapedValue, key))
		}
		sb.WriteString("</PARAMETERS>")
	}
	// Create URL without query parameters
	u := r.serverUrl // Create a copy of the URL

	// Create values for parameters
	reqParams := url.Values{}
	reqParams.Set("requestformat", "JSON")
	reqParams.Set("action", action)
	reqParams.Set("username", r.username)
	reqParams.Set("apiaccesskey", r.apiKey)
	if sb.String() != "" {
		reqParams.Set("parameters", sb.String())
	}

	// Create a new request with POST method, passing reqParams in the body
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(reqParams.Encode()))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Set the content type header for form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Send the request using default client
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	body, err := io.ReadAll(resp.Body)

	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			return
		}
	}(resp.Body)
	var apiData apiResponse
	err = json.Unmarshal(body, &apiData)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %s", err)
	}
	if apiData.ERROR != nil {
		return nil, fmt.Errorf("error from server: %s", apiData.ERROR)
	}

	return body, nil
}

func (r *repo) ping() bool {
	_, err := r.makeApiRequest("accountinfo", nil)
	if err != nil {
		if strings.Contains(err.Error(), "error from server") {
			return true
		}
	}
	return false
}

func (r *repo) getAccountInfo() (accountDetails, error) {
	jsonData, err := r.makeApiRequest("accountinfo", nil)
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
	jsonData, err := r.makeApiRequest("imeiservicelist", nil)
	if err != nil {
		return nil, err
	}
	var root apiResponse
	err = json.Unmarshal(jsonData, &root)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return root.SUCCESS[0].List, nil
}
