package models

type SmPolicyNotification struct {
	SmPolicyDecision *SmPolicyDecision `json:"smPolicyDecision,omitempty"`
	ResourceUri      string            `json:"resourceUri,omitempty"`
}
