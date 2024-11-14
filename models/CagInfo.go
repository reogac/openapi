package models
type CagInfo struct {
	 AllowedCagList	[]string	`json:"allowedCagList"`
	 CagOnlyIndicator	*bool	`json:"cagOnlyIndicator,omitempty"`
}
