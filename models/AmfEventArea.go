package models

type AmfEventArea struct {
	LadnInfo     *LadnInfo     `json:"ladnInfo,omitempty"`
	SNssai       *Snssai       `json:"sNssai,omitempty"`
	NsiId        string        `json:"nsiId,omitempty"`
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
}
