package models
type AdditionalMeasurement struct {
	 UnexpLoc	*NetworkAreaInfo	`json:"unexpLoc,omitempty"`
	 UnexpFlowTeps	[]IpEthFlowDescription	`json:"unexpFlowTeps,omitempty"`
	 UnexpWakes	[]string	`json:"unexpWakes,omitempty"`
	 DdosAttack	*AddressList	`json:"ddosAttack,omitempty"`
	 WrgDest	*AddressList	`json:"wrgDest,omitempty"`
	 Circums	[]CircumstanceDescription	`json:"circums,omitempty"`
}
