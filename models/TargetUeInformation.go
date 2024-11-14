package models
type TargetUeInformation struct {
	 IntGroupIds	[]string	`json:"intGroupIds,omitempty"`
	 AnyUe	*bool	`json:"anyUe,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
	 Gpsis	[]string	`json:"gpsis,omitempty"`
}
