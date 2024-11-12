package models

type RoutingAreaId struct {
	Lac    string `json:"lac"`
	Rac    string `json:"rac"`
	PlmnId PlmnId `json:"plmnId"`
}
