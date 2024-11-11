type UpSecurityInfo struct {
	 UpSecurity	UpSecurity	`json:"upSecurity"`
	 MaxIntegrityProtectedDataRateUl	string	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 MaxIntegrityProtectedDataRateDl	string	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
	 SecurityResult	*SecurityResult	`json:"securityResult,omitempty"`
}
