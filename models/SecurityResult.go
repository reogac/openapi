package models

type SecurityResult struct {
	IntegrityProtectionResult       string `json:"integrityProtectionResult,omitempty"`
	ConfidentialityProtectionResult string `json:"confidentialityProtectionResult,omitempty"`
}
