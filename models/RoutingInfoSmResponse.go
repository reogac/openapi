package models

type RoutingInfoSmResponse struct {
	IpSmGw      *IpSmGwInfo       `json:"ipSmGw,omitempty"`
	SmsRouter   *SmsRouterInfo    `json:"smsRouter,omitempty"`
	Supi        string            `json:"supi,omitempty"`
	Smsf3Gpp    *SmsfRegistration `json:"smsf3Gpp,omitempty"`
	SmsfNon3Gpp *SmsfRegistration `json:"smsfNon3Gpp,omitempty"`
}