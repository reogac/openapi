package models

type IpSmGwRegistration struct {
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
}
