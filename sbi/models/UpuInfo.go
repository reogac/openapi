/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuInfo struct {
	UpuRegInd               *bool     `json:"upuRegInd,omitempty"`
	UpuAckInd               *bool     `json:"upuAckInd,omitempty"`
	UpuMacIausf             string    `json:"upuMacIausf,omitempty"`
	CounterUpu              string    `json:"counterUpu,omitempty"`
	ProvisioningTime        string    `json:"provisioningTime"`
	UpuTransparentContainer string    `json:"upuTransparentContainer,omitempty"`
	UpuDataList             []UpuData `json:"upuDataList,omitempty"`
}