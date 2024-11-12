package models

type EcsServerAddr struct {
	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty"`
	EcsUriList       []string `json:"ecsUriList,omitempty"`
	EcsProviderId    string   `json:"ecsProviderId,omitempty"`
	EcsFqdnList      []string `json:"ecsFqdnList,omitempty"`
}
