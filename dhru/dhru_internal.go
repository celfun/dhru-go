package dhru

// apiResponse contains SUCCESS array and API version
type apiResponse struct {
	SUCCESS    []successItem `json:"SUCCESS"`
	ERROR      []errorItem   `json:"ERROR"`
	APIVersion string        `json:"apiversion"`
}

// errorItem represents an error message
type errorItem struct {
	Message string `json:"MESSAGE"`
}

// successItem contains account information, message, and list of groups
type successItem struct {
	AccountInfo accountDetails `json:"AccountInfo"`
	//AccoutInfo  accountDetails `json:"AccoutInfo"` // Note: keeping the typo as in original JSON
	Message string `json:"MESSAGE"`
	//List    dhruList `json:"LIST"`
	List map[string]group `json:"LIST"` // Direct map instead of RawMessage
}

// accountDetails contains account-related information
type accountDetails struct {
	Credit    string `json:"credit"`
	CreditRaw string `json:"creditraw"`
	Currency  string `json:"currency"`
	Mail      string `json:"mail"`
}

// dhruList contains a map of groups
type dhruList struct {
	Groups map[string]group `json:"-"`
}

// group represents a group with name, type, and associated services
type group struct {
	GroupName string                 `json:"GROUPNAME"`
	GroupType string                 `json:"GROUPTYPE"`
	Services  map[string]dhruService `json:"SERVICES"`
}

// dhruService contains details about a service, including custom configurations
type dhruService struct {
	ServiceID         string        `json:"SERVICEID"`
	ServiceType       string        `json:"SERVICETYPE"`
	Server            string        `json:"SERVER"`
	MinQnt            int           `json:"MINQNT"`
	MaxQnt            int           `json:"MAXQNT"`
	Custom            custom        `json:"CUSTOM"`
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
	RequiresCustom    []customField `json:"Requires.custom"`
}

// custom contains custom configuration for a service
type custom struct {
	Allow      string `json:"allow"`
	Bulk       string `json:"bulk"`
	CustomName string `json:"customname"`
	CustomInfo string `json:"custominfo"`
	CustomLen  any    `json:"customlen"`
	MaxLength  any    `json:"maxlength"`
	Regex      string `json:"regex"`
	IsAlpha    string `json:"isalpha"`
}

// customField represents a custom field with its attributes
type customField struct {
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
