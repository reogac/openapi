package models
type SecurityResult struct {
	 IntegrityProtectionResult	ProtectionResult	`json:"integrityProtectionResult,omitempty"`
	 ConfidentialityProtectionResult	ProtectionResult	`json:"confidentialityProtectionResult,omitempty"`
}
