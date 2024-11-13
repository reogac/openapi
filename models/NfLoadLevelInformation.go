package models

type NfLoadLevelInformation struct {
	NfType             NFType    `json:"nfType,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
}
