package models
type UeContextInSmfDataSubFilter struct {
	 DnnList	[]string	`json:"dnnList,omitempty"`
	 SnssaiList	[]Snssai	`json:"snssaiList,omitempty"`
	 EmergencyInd	*bool	`json:"emergencyInd,omitempty"`
}
