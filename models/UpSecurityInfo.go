package models

type UpSecurityInfo struct {
	SecurityResult                  *SecurityResult `json:"securityResult,omitempty"`
	UpSecurity                      UpSecurity      `json:"upSecurity"`
	MaxIntegrityProtectedDataRateUl string          `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl string          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
}
