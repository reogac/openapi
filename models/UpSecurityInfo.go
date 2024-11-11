package models

type UpSecurityInfo struct {
	MaxIntegrityProtectedDataRateDl string          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SecurityResult                  *SecurityResult `json:"securityResult,omitempty"`
	UpSecurity                      UpSecurity      `json:"upSecurity"`
	MaxIntegrityProtectedDataRateUl string          `json:"maxIntegrityProtectedDataRateUl,omitempty"`
}
