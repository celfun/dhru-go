package dhru

import (
	"net/url"
	"strconv"
)

type Session struct {
	repo *repo
}

func NewDhruSession(inputURL string, username string, apiKey string) (*Session, error) {
	parsedUrl, err := url.Parse(inputURL)
	if err != nil {
		return nil, err
	}
	return &Session{
		repo: newRepo(parsedUrl, username, apiKey),
	}, nil
}

func (s *Session) GetAccountInfo() (*AccountInfo, error) {
	accountInfoData, err := s.repo.getAccountInfo()
	if err != nil {
		return nil, err
	}
	floatCredits, err := strconv.ParseFloat(accountInfoData.CreditRaw, 32)
	if err != nil {
		return nil, err
	}
	return &AccountInfo{
		Credits:  floatCredits,
		Currency: accountInfoData.Currency,
		Email:    accountInfoData.Mail,
	}, nil
}

func (s *Session) GetImeiServiceList() ([]Service, error) {
	imeiServiceListData, err := s.repo.getImeiServiceList()
	if err != nil {
		return nil, err
	}
	var flatServices []Service

	for groupName, groupData := range imeiServiceListData {
		// Process each service in the group
		for _, service := range groupData.Services {
			floatCredit, err2 := strconv.ParseFloat(service.Credit, 64)
			if err2 != nil {
				return nil, err2
			}
			intServiceID, err2 := strconv.ParseInt(service.ServiceID, 10, 64)
			if err2 != nil {
				return nil, err2
			}
			flatService := Service{
				GroupName:         groupName,
				GroupType:         service.ServiceType,
				ServiceID:         intServiceID,
				ServiceName:       service.ServiceName,
				ServiceType:       service.ServiceType,
				Server:            service.Server,
				MinQnt:            service.MinQnt,
				MaxQnt:            service.MaxQnt,
				Credit:            floatCredit,
				Time:              service.Time,
				Info:              service.Info,
				RequiresNetwork:   service.RequiresNetwork,
				RequiresMobile:    service.RequiresMobile,
				RequiresProvider:  service.RequiresProvider,
				RequiresPIN:       service.RequiresPIN,
				RequiresKBH:       service.RequiresKBH,
				RequiresMEP:       service.RequiresMEP,
				RequiresPRD:       service.RequiresPRD,
				RequiresType:      service.RequiresType,
				RequiresLocks:     service.RequiresLocks,
				RequiresReference: service.RequiresReference,
				RequiresSN:        service.RequiresSN,
				RequiresSecRO:     service.RequiresSecRO,
				CustomAllow:       service.Custom.Allow,
				CustomBulk:        service.Custom.Bulk,
				CustomName:        service.Custom.CustomName,
				CustomInfo:        service.Custom.CustomInfo,
				CustomLen:         service.Custom.CustomLen,
				MaxLength:         service.Custom.MaxLength,
				Regex:             service.Custom.Regex,
				IsAlpha:           service.Custom.IsAlpha,
				RequiresCustom:    service.RequiresCustom,
			}
			flatServices = append(flatServices, flatService)
		}
	}

	return flatServices, nil
}
