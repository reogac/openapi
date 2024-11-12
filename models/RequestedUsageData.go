package models

type RequestedUsageData struct {
	RefUmIds []string `json:"refUmIds,omitempty"`
	AllUmIds *bool    `json:"allUmIds,omitempty"`
}
