package models

type N2SmInformation struct {
	PduSessionId   int            `json:"pduSessionId"`
	N2InfoContent  *N2InfoContent `json:"n2InfoContent,omitempty"`
	SNssai         *Snssai        `json:"sNssai,omitempty"`
	HomePlmnSnssai *Snssai        `json:"homePlmnSnssai,omitempty"`
	IwkSnssai      *Snssai        `json:"iwkSnssai,omitempty"`
	SubjectToHo    *bool          `json:"subjectToHo,omitempty"`
}
