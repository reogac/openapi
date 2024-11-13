package models

type IpSmGwRegistration struct {
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
}
