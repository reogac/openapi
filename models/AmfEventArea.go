package models
type AmfEventArea struct {
	 PresenceInfo	*PresenceInfo	`json:"presenceInfo,omitempty"`
	 LadnInfo	*LadnInfo	`json:"ladnInfo,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 NsiId	string	`json:"nsiId,omitempty"`
}
