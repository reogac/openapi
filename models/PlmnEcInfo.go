package models
type PlmnEcInfo struct {
	 PlmnId	PlmnId	`json:"plmnId"`
	 EcRestrictionDataWb	*EcRestrictionDataWb	`json:"ecRestrictionDataWb,omitempty"`
	 EcRestrictionDataNb	*bool	`json:"ecRestrictionDataNb,omitempty"`
}
