/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RgAuthenticationInfo struct {
	AuthenticatedInd  bool   `json:"authenticatedInd"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	Suci              string `json:"suci"`
}
