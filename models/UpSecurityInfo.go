package models
type UpSecurityInfo struct {
	 SecurityResult	*SecurityResult	`json:"securityResult,omitempty"`
	 UpSecurity	UpSecurity	`json:"upSecurity"`
	 MaxIntegrityProtectedDataRateUl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 MaxIntegrityProtectedDataRateDl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
}
