package dhru

import (
	"encoding/json"
)

// UnmarshalJSON Implement UnmarshalJSON to handle the dynamic group names
func (l *dhruList) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	l.Groups = make(map[string]group)
	for groupName, groupData := range raw {
		var group group
		if err := json.Unmarshal(groupData, &group); err != nil {
			return err
		}
		l.Groups[groupName] = group
	}
	return nil
}
