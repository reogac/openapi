package models

type MbsSessionId struct {
	Nid  string `json:"nid,omitempty"`
	Tmgi *Tmgi  `json:"tmgi,omitempty"`
	Ssm  *Ssm   `json:"ssm,omitempty"`
}
