package dhru

import "net/url"

type Dhru struct {
	repo *Repo
}

func NewDhruSession(domain string, apiKey string) *Dhru {
	serverUrl, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}
	return &Dhru{
		repo: NewRepo(serverUrl, apiKey),
	}
}

func (s *Dhru) GetAccountInfo() map[string]string {
	//accountInfoData := s.repo.getAccountInfo()
	accountInfo := make(map[string]string)
	return accountInfo
}
