package models
type RoutingAreaId struct {
	 Rac	string	`json:"rac"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 Lac	string	`json:"lac"`
}
