/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IpSmGwRegistration struct {
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
}
