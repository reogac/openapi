package models
type SmfSelectionData struct {
	 UnsuppDnn	*bool	`json:"unsuppDnn,omitempty"`
	 Candidates	map[string]CandidateForReplacement	`json:"candidates,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 MappingSnssai	*Snssai	`json:"mappingSnssai,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
}
