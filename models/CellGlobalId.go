package models
type CellGlobalId struct {
	 CellId	string	`json:"cellId"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 Lac	string	`json:"lac"`
}
