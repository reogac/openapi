package models

type NfLoadLevelInformation struct {
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
}
