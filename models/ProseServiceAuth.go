package models
type ProseServiceAuth struct {
	 ProseL3RelayAuth	string	`json:"proseL3RelayAuth,omitempty"`
	 ProseL2RemoteAuth	string	`json:"proseL2RemoteAuth,omitempty"`
	 ProseL3RemoteAuth	string	`json:"proseL3RemoteAuth,omitempty"`
	 ProseDirectDiscoveryAuth	string	`json:"proseDirectDiscoveryAuth,omitempty"`
	 ProseDirectCommunicationAuth	string	`json:"proseDirectCommunicationAuth,omitempty"`
	 ProseL2RelayAuth	string	`json:"proseL2RelayAuth,omitempty"`
}
