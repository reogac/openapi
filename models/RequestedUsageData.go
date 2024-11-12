package models

type RequestedUsageData struct {
	AllUmIds *bool    `json:"allUmIds,omitempty"`
	RefUmIds []string `json:"refUmIds,omitempty"`
}
