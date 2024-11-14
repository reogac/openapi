package models
type ChangeItem struct {
	 Path	string	`json:"path"`
	 From	string	`json:"from,omitempty"`
	 Op	ChangeType	`json:"op"`
}
