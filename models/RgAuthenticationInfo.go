package models
type RgAuthenticationInfo struct {
	 Suci	string	`json:"suci"`
	 AuthenticatedInd	bool	`json:"authenticatedInd"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
