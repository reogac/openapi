package models
type TsnBridgeInfo struct {
	 BridgeId	*int	`json:"bridgeId,omitempty"`
	 DsttAddr	string	`json:"dsttAddr,omitempty"`
	 DsttPortNum	*int	`json:"dsttPortNum,omitempty"`
	 DsttResidTime	*int	`json:"dsttResidTime,omitempty"`
}
