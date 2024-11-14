package models
type NfLoadLevelInformation struct {
	 NfMemoryUsage	*int	`json:"nfMemoryUsage,omitempty"`
	 NfStorageUsage	*int	`json:"nfStorageUsage,omitempty"`
	 NfLoadLevelAverage	*int	`json:"nfLoadLevelAverage,omitempty"`
	 NfLoadLevelpeak	*int	`json:"nfLoadLevelpeak,omitempty"`
	 NfLoadAvgInAoi	*int	`json:"nfLoadAvgInAoi,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 NfStatus	*NfStatus	`json:"nfStatus,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId,omitempty"`
	 NfSetId	string	`json:"nfSetId,omitempty"`
	 NfCpuUsage	*int	`json:"nfCpuUsage,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 NfType	NFType	`json:"nfType,omitempty"`
}
