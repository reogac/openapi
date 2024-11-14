package models
type SmPolicyNotification struct {
	 ResourceUri	string	`json:"resourceUri,omitempty"`
	 SmPolicyDecision	*SmPolicyDecision	`json:"smPolicyDecision,omitempty"`
}
