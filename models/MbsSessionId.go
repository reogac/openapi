package models

type MbsSessionId struct {
	Ssm  *Ssm   `json:"ssm,omitempty"`
	Nid  string `json:"nid,omitempty"`
	Tmgi *Tmgi  `json:"tmgi,omitempty"`
}
