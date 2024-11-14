package models
type TargetUeInformation struct {
	 AnyUe	*bool	`json:"anyUe,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
	 Gpsis	[]string	`json:"gpsis,omitempty"`
	 IntGroupIds	[]string	`json:"intGroupIds,omitempty"`
}
