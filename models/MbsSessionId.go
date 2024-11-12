package models

type MbsSessionId struct {
	Tmgi *Tmgi  `json:"tmgi,omitempty"`
	Ssm  *Ssm   `json:"ssm,omitempty"`
	Nid  string `json:"nid,omitempty"`
}
