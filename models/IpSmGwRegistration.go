package models

type IpSmGwRegistration struct {
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
}
