package models
type AmfEventArea struct {
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 NsiId	string	`json:"nsiId,omitempty"`
	 PresenceInfo	*PresenceInfo	`json:"presenceInfo,omitempty"`
	 LadnInfo	*LadnInfo	`json:"ladnInfo,omitempty"`
}
