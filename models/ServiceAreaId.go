package models

type ServiceAreaId struct {
	Lac    string `json:"lac"`
	Sac    string `json:"sac"`
	PlmnId PlmnId `json:"plmnId"`
}
