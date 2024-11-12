package models

type SmPolicyControl struct {
	Policy  SmPolicyDecision    `json:"policy"`
	Context SmPolicyContextData `json:"context"`
}
