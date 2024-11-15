/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NfLoadLevelInformation struct {
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
}
