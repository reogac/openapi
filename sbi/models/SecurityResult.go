/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecurityResult struct {
	IntegrityProtectionResult       ProtectionResult `json:"integrityProtectionResult,omitempty"`
	ConfidentialityProtectionResult ProtectionResult `json:"confidentialityProtectionResult,omitempty"`
}
