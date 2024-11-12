package models

type NfLoadLevelInformation struct {
	NfType             NFType    `json:"nfType,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
}
