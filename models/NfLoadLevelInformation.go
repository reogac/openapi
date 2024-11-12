package models

type NfLoadLevelInformation struct {
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	NfType             string    `json:"nfType,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
}
