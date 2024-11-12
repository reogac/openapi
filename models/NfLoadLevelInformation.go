package models

type NfLoadLevelInformation struct {
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
}
