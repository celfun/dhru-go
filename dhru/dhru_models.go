package dhru

// AccountInfo contains user account information
type AccountInfo struct {
	Email    string  `json:"mail"`
	Credits  float64 `json:"credit"`
	Currency string  `json:"currency"`
}

// Service represents the flattened service structure
type Service struct {
	GroupName         string        `json:"group_name"`
	GroupType         string        `json:"group_type"`
	ServiceID         int64         `json:"service_id"`
	ServiceName       string        `json:"service_name"`
	ServiceType       string        `json:"service_type"`
	Server            string        `json:"server"`
	MinQnt            int           `json:"min_qnt"`
	MaxQnt            int           `json:"max_qnt"`
	Credit            float64       `json:"credit"`
	Time              string        `json:"time"`
	Info              string        `json:"info"`
	RequiresNetwork   string        `json:"requires_network"`
	RequiresMobile    string        `json:"requires_mobile"`
	RequiresProvider  string        `json:"requires_provider"`
	RequiresPIN       string        `json:"requires_pin"`
	RequiresKBH       string        `json:"requires_kbh"`
	RequiresMEP       string        `json:"requires_mep"`
	RequiresPRD       string        `json:"requires_prd"`
	RequiresType      string        `json:"requires_type"`
	RequiresLocks     string        `json:"requires_locks"`
	RequiresReference string        `json:"requires_reference"`
	RequiresSN        string        `json:"requires_sn"`
	RequiresSecRO     string        `json:"requires_secro"`
	CustomAllow       string        `json:"custom_allow"`
	CustomBulk        string        `json:"custom_bulk"`
	CustomName        string        `json:"custom_name"`
	CustomInfo        string        `json:"custom_info"`
	CustomLen         any           `json:"custom_len"`
	MaxLength         any           `json:"max_length"`
	Regex             string        `json:"regex"`
	IsAlpha           string        `json:"is_alpha"`
	RequiresCustom    []customField `json:"requires_custom"`
}
