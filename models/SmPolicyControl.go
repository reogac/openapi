package models

type SmPolicyControl struct {
	Context SmPolicyContextData `json:"context"`
	Policy  SmPolicyDecision    `json:"policy"`
}
