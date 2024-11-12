package models

type RgAuthenticationInfo struct {
	AuthenticatedInd  bool   `json:"authenticatedInd"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	Suci              string `json:"suci"`
}
