package models
type RgAuthenticationInfo struct {
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 Suci	string	`json:"suci"`
	 AuthenticatedInd	bool	`json:"authenticatedInd"`
}
