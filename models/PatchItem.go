package models
type PatchItem struct {
	 Path	string	`json:"path"`
	 From	string	`json:"from,omitempty"`
	 Op	PatchOperation	`json:"op"`
}
