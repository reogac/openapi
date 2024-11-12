package models

type SecurityResult struct {
	ConfidentialityProtectionResult string `json:"confidentialityProtectionResult,omitempty"`
	IntegrityProtectionResult       string `json:"integrityProtectionResult,omitempty"`
}
