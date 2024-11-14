package models
type ChangeItem struct {
	 From	string	`json:"from,omitempty"`
	 Op	ChangeType	`json:"op"`
	 Path	string	`json:"path"`
}
