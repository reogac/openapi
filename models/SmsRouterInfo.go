package models

type SmsRouterInfo struct {
	NfInstanceId    string                      `json:"nfInstanceId,omitempty"`
	DiameterAddress *NetworkNodeDiameterAddress `json:"diameterAddress,omitempty"`
	MapAddress      string                      `json:"mapAddress,omitempty"`
	RouterIpv4      string                      `json:"routerIpv4,omitempty"`
	RouterIpv6      string                      `json:"routerIpv6,omitempty"`
	RouterFqdn      string                      `json:"routerFqdn,omitempty"`
}
