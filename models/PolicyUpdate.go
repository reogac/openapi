package models
type PolicyUpdate struct {
	 Pras	map[string]PresenceInfo	`json:"pras,omitempty"`
	 ResourceUri	string	`json:"resourceUri"`
	 UePolicy	string	`json:"uePolicy,omitempty"`
	 N2Pc5Pol	*N2InfoContent	`json:"n2Pc5Pol,omitempty"`
	 N2Pc5ProSePol	*N2InfoContent	`json:"n2Pc5ProSePol,omitempty"`
	 Triggers	[]string	`json:"triggers,omitempty"`
}
