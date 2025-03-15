package dhru

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON Implement UnmarshalJSON to handle the dynamic group names
func (l *List) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	l.Groups = make(map[string]Group)
	for groupName, groupData := range raw {
		var group Group
		if err := json.Unmarshal(groupData, &group); err != nil {
			return err
		}
		l.Groups[groupName] = group
	}
	return nil
}

// FlattenIMEIServiceList flattens the nested JSON structure into a list of services
func flattenIMEIServiceList(jsonData []byte) ([]FlatService, error) {
	var root RootObject
	if err := json.Unmarshal(jsonData, &root); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	var flatServices []FlatService

	// Process each success item
	for _, successItem := range root.IMEIServiceList.Success {
		// Process each group
		for groupName, group := range successItem.List.Groups {
			// Process each service in the group
			for serviceID, service := range group.Services {
				flatService := FlatService{
					GroupName:         groupName,
					GroupType:         group.GroupType,
					ServiceID:         serviceID,
					ServiceName:       service.ServiceName,
					ServiceType:       service.ServiceType,
					Server:            service.Server,
					MinQnt:            service.MinQnt,
					MaxQnt:            service.MaxQnt,
					Credit:            service.Credit,
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
	}

	return flatServices, nil
}

func mapAccountInfo(jsonData []byte) (*AccountDetails, error) {
	var root AccountInfoResponse
	if err := json.Unmarshal(jsonData, &root); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return &root.ACCOUNTINFO.SUCCESS[0].AccoutInfo, nil
}
