package dhru

//Account info struct

// AccountInfoResponse Response represents the top-level JSON structure
type AccountInfoResponse struct {
	ACCOUNTINFO AccountInfo `json:"ACCOUNTINFO"`
}

// AccountInfo contains SUCCESS array and API version
type AccountInfo struct {
	SUCCESS    []SuccessItem `json:"SUCCESS"`
	APIVersion string        `json:"apiversion"`
}

// AccountDetails contains user account information
type AccountDetails struct {
	Credit    string  `json:"credit"`
	CreditRaw float64 `json:"creditraw"`
	Currency  string  `json:"currency"`
	Mail      string  `json:"mail"`
}

// Service Structs

// CustomField Define structs to match the input JSON structure
type CustomField struct {
	Type         string `json:"type"`
	FieldName    string `json:"fieldname"`
	FieldType    string `json:"fieldtype"`
	Description  string `json:"description"`
	Required     string `json:"required"`
	FieldOptions string `json:"fieldoptions"`
	RegExpr      string `json:"regexpr"`
	AdminOnly    string `json:"adminonly"`
	Enc          string `json:"enc"`
}

type Custom struct {
	Allow      string `json:"allow"`
	Bulk       string `json:"bulk"`
	CustomName string `json:"customname"`
	CustomInfo string `json:"custominfo"`
	CustomLen  any    `json:"customlen"`
	MaxLength  any    `json:"maxlength"`
	Regex      string `json:"regex"`
	IsAlpha    string `json:"isalpha"`
}

type Service struct {
	ServiceID         string        `json:"SERVICEID"`
	ServiceType       string        `json:"SERVICETYPE"`
	Server            string        `json:"SERVER"`
	MinQnt            int           `json:"MINQNT"`
	MaxQnt            int           `json:"MAXQNT"`
	Custom            Custom        `json:"CUSTOM"`
	ServiceName       string        `json:"SERVICENAME"`
	Credit            string        `json:"CREDIT"`
	Time              string        `json:"TIME"`
	Info              string        `json:"INFO"`
	RequiresNetwork   string        `json:"Requires.Network"`
	RequiresMobile    string        `json:"Requires.Mobile"`
	RequiresProvider  string        `json:"Requires.Provider"`
	RequiresPIN       string        `json:"Requires.PIN"`
	RequiresKBH       string        `json:"Requires.KBH"`
	RequiresMEP       string        `json:"Requires.MEP"`
	RequiresPRD       string        `json:"Requires.PRD"`
	RequiresType      string        `json:"Requires.Type"`
	RequiresLocks     string        `json:"Requires.Locks"`
	RequiresReference string        `json:"Requires.Reference"`
	RequiresSN        string        `json:"Requires.SN"`
	RequiresSecRO     string        `json:"Requires.SecRO"`
	RequiresCustom    []CustomField `json:"Requires.Custom"`
}

type Group struct {
	GroupName string             `json:"GROUPNAME"`
	GroupType string             `json:"GROUPTYPE"`
	Services  map[string]Service `json:"SERVICES"`
}

type List struct {
	Groups map[string]Group `json:"-"`
}

type SuccessItem struct {
	AccountInfo AccountDetails `json:"AccountInfo"`
	AccoutInfo  AccountDetails `json:"AccoutInfo"` // Note: keeping the typo as in original JSON
	Message     string         `json:"MESSAGE"`
	List        List           `json:"LIST"`
}

type IMEIServiceList struct {
	Success []SuccessItem `json:"SUCCESS"`
}

type RootObject struct {
	IMEIServiceList IMEIServiceList `json:"IMEISERVICELIST"`
}

// FlatService represents the flattened service structure
type FlatService struct {
	GroupName         string        `json:"group_name"`
	GroupType         string        `json:"group_type"`
	ServiceID         string        `json:"service_id"`
	ServiceName       string        `json:"service_name"`
	ServiceType       string        `json:"service_type"`
	Server            string        `json:"server"`
	MinQnt            int           `json:"min_qnt"`
	MaxQnt            int           `json:"max_qnt"`
	Credit            string        `json:"credit"`
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
	RequiresCustom    []CustomField `json:"requires_custom"`
}
