package models

type SmfSelectionData struct {
	Dnn           string                             `json:"dnn,omitempty"`
	UnsuppDnn     *bool                              `json:"unsuppDnn,omitempty"`
	Candidates    map[string]CandidateForReplacement `json:"candidates,omitempty"`
	Snssai        *Snssai                            `json:"snssai,omitempty"`
	MappingSnssai *Snssai                            `json:"mappingSnssai,omitempty"`
}
