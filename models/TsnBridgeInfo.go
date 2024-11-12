package models

type TsnBridgeInfo struct {
	DsttPortNum   *int   `json:"dsttPortNum,omitempty"`
	DsttResidTime *int   `json:"dsttResidTime,omitempty"`
	BridgeId      *int   `json:"bridgeId,omitempty"`
	DsttAddr      string `json:"dsttAddr,omitempty"`
}
