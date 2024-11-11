package models
type Ecgi struct {
	 Nid	string	`json:"nid,omitempty"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 EutraCellId	string	`json:"eutraCellId"`
}
