package models
type DispersionArea struct {
	 TaiList	[]Tai	`json:"taiList,omitempty"`
	 NcgiList	[]Ncgi	`json:"ncgiList,omitempty"`
	 EcgiList	[]Ecgi	`json:"ecgiList,omitempty"`
	 N3gaInd	*bool	`json:"n3gaInd,omitempty"`
}
