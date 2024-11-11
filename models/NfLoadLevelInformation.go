package models

type NfLoadLevelInformation struct {
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfType             string    `json:"nfType,omitempty"`
}
