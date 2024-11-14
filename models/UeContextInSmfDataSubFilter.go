package models
type UeContextInSmfDataSubFilter struct {
	 EmergencyInd	*bool	`json:"emergencyInd,omitempty"`
	 DnnList	[]string	`json:"dnnList,omitempty"`
	 SnssaiList	[]Snssai	`json:"snssaiList,omitempty"`
}
