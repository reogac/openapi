package models
type SteeringInfo struct {
	 PlmnId	PlmnId	`json:"plmnId"`
	 AccessTechList	[]string	`json:"accessTechList,omitempty"`
}
