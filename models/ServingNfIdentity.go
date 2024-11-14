package models
type ServingNfIdentity struct {
	 Guami	*Guami	`json:"guami,omitempty"`
	 AnGwAddr	*AnGwAddress	`json:"anGwAddr,omitempty"`
	 SgsnAddr	*SgsnAddress	`json:"sgsnAddr,omitempty"`
	 ServNfInstId	string	`json:"servNfInstId,omitempty"`
}
