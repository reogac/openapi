/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationTransferReqData struct {
	N2Information     N2InfoContainer   `json:"n2Information"`
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
	TaiList           []Tai             `json:"taiList,omitempty"`
	RatSelector       RatSelector       `json:"ratSelector,omitempty"`
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
}
