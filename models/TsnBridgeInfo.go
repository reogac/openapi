package models

type TsnBridgeInfo struct {
	DsttResidTime *int   `json:"dsttResidTime,omitempty"`
	BridgeId      *int   `json:"bridgeId,omitempty"`
	DsttAddr      string `json:"dsttAddr,omitempty"`
	DsttPortNum   *int   `json:"dsttPortNum,omitempty"`
}
