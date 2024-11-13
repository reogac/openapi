package models

type SecurityResult struct {
	ConfidentialityProtectionResult ProtectionResult `json:"confidentialityProtectionResult,omitempty"`
	IntegrityProtectionResult       ProtectionResult `json:"integrityProtectionResult,omitempty"`
}
