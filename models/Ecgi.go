package models
type Ecgi struct {
	 PlmnId	PlmnId	`json:"plmnId"`
	 EutraCellId	string	`json:"eutraCellId"`
	 Nid	string	`json:"nid,omitempty"`
}
