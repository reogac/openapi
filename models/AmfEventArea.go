package models

type AmfEventArea struct {
	NsiId        string        `json:"nsiId,omitempty"`
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
	LadnInfo     *LadnInfo     `json:"ladnInfo,omitempty"`
	SNssai       *Snssai       `json:"sNssai,omitempty"`
}
