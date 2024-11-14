package models
type NfLoadLevelInformation struct {
	 NfCpuUsage	*int	`json:"nfCpuUsage,omitempty"`
	 NfMemoryUsage	*int	`json:"nfMemoryUsage,omitempty"`
	 NfStorageUsage	*int	`json:"nfStorageUsage,omitempty"`
	 NfLoadAvgInAoi	*int	`json:"nfLoadAvgInAoi,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 NfType	NFType	`json:"nfType,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId,omitempty"`
	 NfSetId	string	`json:"nfSetId,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 NfStatus	*NfStatus	`json:"nfStatus,omitempty"`
	 NfLoadLevelAverage	*int	`json:"nfLoadLevelAverage,omitempty"`
	 NfLoadLevelpeak	*int	`json:"nfLoadLevelpeak,omitempty"`
}
