package models
type ServingNfIdentity struct {
	 ServNfInstId	string	`json:"servNfInstId,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 AnGwAddr	*AnGwAddress	`json:"anGwAddr,omitempty"`
	 SgsnAddr	*SgsnAddress	`json:"sgsnAddr,omitempty"`
}
