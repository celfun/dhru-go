package dhru

import "net/url"

type Repo struct {
	serverUrl url.URL
	apiKey    string
}

func NewRepo(serverUrl *url.URL, apiKey string) *Repo {
	return &Repo{
		serverUrl: *serverUrl,
		apiKey:    apiKey,
	}
}

func (r *Repo) getAccountInfo() {

}
