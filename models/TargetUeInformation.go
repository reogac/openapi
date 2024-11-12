package models

type TargetUeInformation struct {
	Gpsis       []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
	AnyUe       *bool    `json:"anyUe,omitempty"`
	Supis       []string `json:"supis,omitempty"`
}
