/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NfLoadLevelInformation struct {
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
}
