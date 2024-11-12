package models

type IpSmGwRegistration struct {
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
}
