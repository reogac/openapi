package models
type TsnBridgeInfo struct {
	 DsttAddr	string	`json:"dsttAddr,omitempty"`
	 DsttPortNum	*int	`json:"dsttPortNum,omitempty"`
	 DsttResidTime	*int	`json:"dsttResidTime,omitempty"`
	 BridgeId	*int	`json:"bridgeId,omitempty"`
}
