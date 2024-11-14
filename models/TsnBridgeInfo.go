/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TsnBridgeInfo struct {
	DsttResidTime *int   `json:"dsttResidTime,omitempty"`
	BridgeId      *int   `json:"bridgeId,omitempty"`
	DsttAddr      string `json:"dsttAddr,omitempty"`
	DsttPortNum   *int   `json:"dsttPortNum,omitempty"`
}
