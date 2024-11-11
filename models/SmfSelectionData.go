package models
type SmfSelectionData struct {
	 Dnn	string	`json:"dnn,omitempty"`
	 UnsuppDnn	*bool	`json:"unsuppDnn,omitempty"`
	 Candidates	*candidates	`json:"candidates,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 MappingSnssai	*Snssai	`json:"mappingSnssai,omitempty"`
}
