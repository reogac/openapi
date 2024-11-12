package models

type PolicyAssociation struct {
	Triggers      []string                  `json:"triggers,omitempty"`
	Pras          map[string]PresenceInfo   `json:"pras,omitempty"`
	SuppFeat      string                    `json:"suppFeat"`
	Request       *PolicyAssociationRequest `json:"request,omitempty"`
	UePolicy      string                    `json:"uePolicy,omitempty"`
	N2Pc5Pol      *N2InfoContent            `json:"n2Pc5Pol,omitempty"`
	N2Pc5ProSePol *N2InfoContent            `json:"n2Pc5ProSePol,omitempty"`
}
